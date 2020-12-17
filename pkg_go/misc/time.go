package misc

import (
	"fmt"
	"time"
)

/*
  ceil time, e.g. TimeCeil('2020-12-01T17:39:07.123+08:00', "M") -> '2020-12-01T17:00:00+08:00'
    valid unit(key or value)
    H: hour, M: minute, S: second
    y: year, s: season, m: month, w: week, d: day
*/
func TimeCeil(at time.Time, tu string) (out time.Time, err error) {
	var (
		location        *time.Location
		year, day, hour int
		minute, second  int
		month           time.Month
	)

	location = at.Location()
	year, month, day = at.Date()
	hour, minute, second = at.Clock()

	switch tu {
	case "second", "S":
		out = time.Date(year, month, day, hour, minute, second, 0, location)
		if at.Sub(out) > time.Millisecond {
			out = out.Add(time.Second)
		}

	case "minute", "M":
		out = time.Date(year, month, day, hour, minute, 0, 0, location)
		if at.Sub(out) > time.Second {
			out = out.Add(time.Minute)
		}

	case "hour", "H":
		out = time.Date(year, month, day, hour, 0, 0, 0, location)
		if at.Sub(out) > time.Second {
			out = out.Add(time.Hour)
		}

	case "day", "d":
		out = time.Date(year, month, day, 0, 0, 0, 0, location)
		if at.Sub(out) > time.Second {
			out = out.AddDate(0, 0, 1)
		}

	case "week", "w":
		w := int(at.Weekday()) // 0 for Sunday, the first day of a week is Monday
		if w == 0 {
			w = 7
		}

		year, month, day = at.AddDate(0, 0, 1-w).Date()
		out = time.Date(year, month, day, 0, 0, 0, 0, location)
		if at.Sub(out) > time.Second {
			out = out.AddDate(0, 0, 7)
		}

	case "month", "m":
		out = time.Date(year, month, 1, 0, 0, 0, 0, location)
		if at.Sub(out) > time.Second {
			out = out.AddDate(0, 1, 0)
		}

	case "season", "s":
		month = (month-1)/3*3 + 1
		out = time.Date(year, month, 1, 0, 0, 0, 0, location)
		if at.Sub(out) > time.Second {
			out = out.AddDate(0, 3, 0)
		}

	case "year", "y":
		out = time.Date(year, 1, 1, 0, 0, 0, 0, location)
		if at.Sub(out) > time.Second {
			out = out.AddDate(1, 0, 0)
		}
	default:
		err = fmt.Errorf("invalid time unit")
		return
	}

	return
}
