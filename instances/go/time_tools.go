package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DefaultTimeFormat = "2006-01-02T15:04:05-0700"

// Convert between int64 and time, type example s2t, t2u,
//   s for Second, m for Millisecond, u for Microsecond, n for Nanosecond
//   t for time in format 2006-01-02T15:04:05-0700
func TimeConvert(ty, str string) (result string, err error) {
	var (
		n int64
		t time.Time
	)

	dtf := DefaultTimeFormat
	ty = strings.ToLower(ty)

	if strings.HasSuffix(ty, "2t") {
		n, err = strconv.ParseInt(str, 10, 64)
	}

	if strings.HasPrefix(ty, "t2") {
		t, err = time.Parse(DefaultTimeFormat, str)
	}

	if err != nil {
		return
	}

	// fmt.Println(ty, str, n, t)

	switch ty {
	case "s2t":
		t = time.Unix(n, 0)
		result = t.Format(dtf)
	case "m2t":
		t = time.Unix(n/1e3, n%1e3*1e6)
		result = t.Format(dtf)
	case "u2t":
		t = time.Unix(n/1e6, n%1e6*1e3)
		result = t.Format(dtf)
	case "n2t":
		t = time.Unix(n/1e9, n%1e9)
		result = t.Format(dtf)
	case "t2s":
		result = strconv.FormatInt(t.Unix(), 10)
	case "t2m":
		result = strconv.FormatInt(t.UnixNano()/1e6, 10)
	case "t2u":
		result = strconv.FormatInt(t.UnixNano()/1e3, 10)
	case "t2n":
		result = strconv.FormatInt(t.UnixNano(), 10)
	default:
		err = fmt.Errorf("invalid type %s", ty)
		return
	}

	return
}

////
type TimeSplitter struct {
	Start, End time.Time
	Duration   time.Duration
	n, k       int64
}

func NewTimeSplitter(start, end time.Time, d time.Duration) (
	tsp *TimeSplitter, err error) {
	if d < 1 {
		err = fmt.Errorf("invalid time.Duaration")
		return
	}

	if end.Sub(start) == 0 {
		err = fmt.Errorf("start is equal to end")
		return
	}

	tsp = new(TimeSplitter)
	if start.After(end) {
		start, end = end, start
	}

	tsp.Start, tsp.End, tsp.Duration = start, end, d
	tsp.n = int64(end.Sub(start) / d)

	return
}

func (tsp *TimeSplitter) GetNK() (int64, int64) {
	return tsp.n, tsp.k
}

func (tsp *TimeSplitter) Next() (qStart, qEnd time.Time, err error) {
	defer func() { tsp.k++ }()

	if tsp.k > tsp.n {
		err = fmt.Errorf("No more time pieces")
		return
	}

	qStart = tsp.Start.Add(time.Duration(tsp.k) * tsp.Duration)
	if qStart.After(tsp.End) || qStart == tsp.End {
		err = fmt.Errorf("no more time pieces")
		return
	}

	if tsp.k == tsp.n {
		qEnd = tsp.End
	} else {
		qEnd = qStart.Add(tsp.Duration)
	}

	return
}
