package rover

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var invalidTimeUint = errors.New("invalid time unit (not in S:second, " +
	"M:minute, H:hour, d:day, w:week, m:month, s:season, y:year)")

func InvalidTimeUint() (err error) {
	return invalidTimeUint
}

func TimeCeil(from time.Time, tu string) (nt time.Time, err error) {
	var (
		location        *time.Location
		year, day, hour int
		minute, second  int
		month           time.Month
	)

	if tu, err = FixTimeUnit(tu); err != nil {
		return
	}

	location = from.Location()
	year, month, day = from.Date()
	hour, minute, second = from.Clock()

	switch tu {
	case "second":
		nt = time.Date(year, month, day, hour, minute, second, 0, location)
		if from.Sub(nt) > time.Millisecond {
			nt = nt.Add(time.Second)
		}

	case "minute":
		nt = time.Date(year, month, day, hour, minute, 0, 0, location)
		if from.Sub(nt) > time.Second {
			nt = nt.Add(time.Minute)
		}

	case "hour":
		nt = time.Date(year, month, day, hour, 0, 0, 0, location)
		if from.Sub(nt) > time.Second {
			nt = nt.Add(time.Hour)
		}

	case "day":
		nt = time.Date(year, month, day, 0, 0, 0, 0, location)
		if from.Sub(nt) > time.Second {
			nt = nt.AddDate(0, 0, 1)
		}

	case "week":
		w := int(from.Weekday()) // 0 for Sunday, the first day of a week is Monday
		if w == 0 {
			w = 7
		}

		year, month, day = from.AddDate(0, 0, 1-w).Date()
		nt = time.Date(year, month, day, 0, 0, 0, 0, location)
		if from.Sub(nt) > time.Second {
			nt = nt.AddDate(0, 0, 7)
		}

	case "month":
		nt = time.Date(year, month, 1, 0, 0, 0, 0, location)
		if from.Sub(nt) > time.Second {
			nt = nt.AddDate(0, 1, 0)
		}

	case "season":
		month = (month-1)/3*3 + 1
		nt = time.Date(year, month, 1, 0, 0, 0, 0, location)
		if from.Sub(nt) > time.Second {
			nt = nt.AddDate(0, 3, 0)
		}

	case "year":
		nt = time.Date(year, 1, 1, 0, 0, 0, 0, location)
		if from.Sub(nt) > time.Second {
			nt = nt.AddDate(1, 0, 0)
		}
	}

	return
}

func TimeFloor(from time.Time, tu string) (nt time.Time, err error) {
	var (
		location        *time.Location
		year, day, hour int
		second, minute  int
		month           time.Month
	)

	if tu, err = FixTimeUnit(tu); err != nil {
		return
	}

	location = from.Location()
	year, month, day = from.Date()
	hour, minute, second = from.Clock()

	switch tu {
	case "second":
		nt = time.Date(year, month, day, hour, minute, second, 0, location)

	case "minute":
		nt = time.Date(year, month, day, hour, minute, 0, 0, location)

	case "hour":
		nt = time.Date(year, month, day, hour, 0, 0, 0, location)

	case "day":
		nt = time.Date(year, month, day, 0, 0, 0, 0, location)

	case "week":
		w := int(from.Weekday())
		// 0 for Sunday, the first day of a week is Monday
		if w == 0 {
			w = 7
		}

		year, month, day = from.AddDate(0, 0, 1-w).Date()
		nt = time.Date(year, month, day, 0, 0, 0, 0, location)

	case "month":
		nt = time.Date(year, month, 1, 0, 0, 0, 0, location)

	case "season":
		month = (month-1)/3*3 + 1
		nt = time.Date(year, month, 1, 0, 0, 0, 0, location)

	case "year":
		nt = time.Date(year, 1, 1, 0, 0, 0, 0, location)
	}

	return
}

func TimeWalk(from time.Time, n int64, tu string) (nt time.Time, err error) {
	var offset time.Duration

	if tu, err = FixTimeUnit(tu); err != nil {
		return
	}

	switch tu {
	case "second":
		offset = time.Duration(n) * time.Second
		nt = from.Add(offset)

	case "minute":
		offset = time.Duration(n) * time.Minute
		nt = from.Add(offset)

	case "hour":
		offset = time.Duration(n) * time.Hour
		nt = from.Add(offset)

	case "day":
		nt = from.AddDate(0, 0, int(n))

	case "week":
		nt = from.AddDate(0, 0, int(n)*7)

	case "month":
		nt = from.AddDate(0, int(n), 0)

	case "season":
		nt = from.AddDate(0, int(n)*3, 0)

	case "year":
		nt = from.AddDate(int(n), 0, 0)
	}

	return
}

func TimeWalkRange(from time.Time, n int64, tu string) (t0, t1 time.Time, err error) {
	var t time.Time
	if t, err = TimeWalk(from, n, tu); err != nil {
		return
	}

	t0, t1 = from, t
	if from.After(t) {
		t0, t1 = t1, t0
	}

	return
}

func TimeScale(from time.Time, str string, cof func(time.Time, string) (time.Time, error)) (
	nt time.Time, err error) {

	var n int64
	var tu string

	if n, tu, err = ParseNTU(str); err != nil {
		return
	}

	if nt, err = cof(from, tu); err != nil {
		return
	}

	if nt, err = TimeWalk(nt, n, tu); err != nil {
		return
	}

	return
}

func TimeScaleRange(from time.Time, str string, cof func(time.Time, string) (time.Time, error)) (
	t0, t1 time.Time, err error) {

	var t time.Time
	if t, err = TimeScale(from, str, cof); err != nil {
		return
	}

	t0, t1 = from, t
	if from.After(t) {
		t0, t1 = t1, t0
	}

	return
}

func FixTimeUnit(in string) (out string, err error) {
	out = in
	if len(out) == 1 {
		switch out {
		case "S":
			out = "second"
		case "M":
			out = "minute"
		case "H":
			out = "hour"
		case "d":
			out = "day"
		case "w":
			out = "week"
		case "m":
			out = "month"
		case "s":
			out = "season"
		case "y":
			out = "year"
		}
	}

	out = strings.ToLower(out)

	switch out {
	case "second", "minute", "hour", "day", "week", "month", "season", "year":
	default:
		return "", invalidTimeUint
	}

	return
}

func ParseNTU(str string) (n int64, unit string, err error) {
	re, _ := regexp.Compile(`([0-9]+)\s*([a-z]+)`)
	slice := re.FindStringSubmatch(str)

	if len(slice) != 3 {
		err = fmt.Errorf("regular expression not match")
		return
	}

	d, u := slice[1], slice[2]
	if unit, err = FixTimeUnit(u); err != nil {
		return
	}

	n, _ = strconv.ParseInt(d, 10, 64)
	return
}
