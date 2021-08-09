package explore

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"sync/atomic"
	"time"
)

type Pedrail struct {
	Name       string
	Map        map[int64]*int64
	StartAt    time.Time
	n, seconds int
	ticker     *time.Ticker
	quit       chan struct{}
	currentKey func() int64
	do         func(map[int64]int64) error
	infof      func(format string, args ...interface{}) (int, error)
	errorf     func(format string, args ...interface{}) (int, error)
}

// n int:       number time unit for a roll
// dur string:  clock time unit (second/S, minute/M, hour/H)
// name string: name to Pedrail as logging tag
func NewPedrail(n int, dur, name string, do func(map[int64]int64) error) (
	ped *Pedrail, err error) {

	var (
		i   int
		key int64
		at  time.Time
	)

	if n <= 0 {
		return nil, fmt.Errorf("invalid parameter n")
	}

	if at, err = FloorClock(1, dur); err != nil {
		return nil, err
	}

	ped = &Pedrail{StartAt: at, Name: name, n: n, do: do}

	switch dur {
	case "second", "S":
		ped.seconds = 1
		ped.currentKey = func() int64 { return time.Now().Unix() }
	case "minute", "M":
		ped.seconds = 60
		ped.currentKey = func() int64 {
			t := time.Now()
			year, month, day := t.Date()
			hour, minute, _ := t.Clock()
			return time.Date(year, month, day, hour, minute, 0, 0, t.Location()).Unix()
		}
	case "hour", "H":
		ped.seconds = 3600
		ped.currentKey = func() int64 {
			t := time.Now()
			year, month, day := t.Date()
			hour, _, _ := t.Clock()
			return time.Date(year, month, day, hour, 0, 0, 0, t.Location()).Unix()
		}
	}

	ped.ticker = time.NewTicker(time.Duration(ped.n*ped.seconds) * time.Second)
	ped.Map = make(map[int64]*int64, 2*ped.n)
	ped.quit = make(chan struct{})

	for i = 0; i < 2*ped.n; i++ {
		var n int64
		key = ped.StartAt.Add(time.Duration(i*ped.seconds) * time.Second).Unix()
		// fmt.Println(">>>", key)
		ped.Map[key] = &n
	}

	ped.infof = func(format string, args ...interface{}) (int, error) {
		format = fmt.Sprintf("[%s] %s %s\n",
			LogTime("0102T15:04:05"), ped.Name, strings.Trim(format, "\n\r"))

		return fmt.Fprintf(os.Stdout, format, args...)
	}

	ped.errorf = func(format string, args ...interface{}) (int, error) {
		format = fmt.Sprintf("[%s] %s %s\n",
			LogTime("0102T15:04:05"), ped.Name, strings.Trim(format, "\n\r"))

		return fmt.Fprintf(os.Stderr, format, args...)
	}

	return
}

// Pedrail rolls, called by Starr, clear expired key in map and add new keys of comming clock
func (ped *Pedrail) Roll() (mp map[int64]int64) {
	var (
		i      int
		ok     bool
		k, key int64
	)

	key = ped.currentKey()
	for i = 1; i < 2*ped.n; i++ {
		k = time.Unix(key+int64(i)*int64(ped.seconds), 0).Unix()
		if _, ok = ped.Map[k]; !ok {
			var v int64
			ped.Map[k] = &v
		}
	}

	ped.infof("roll: %s", ped.DataJSON())
	mp = make(map[int64]int64, ped.n)
	for k = range ped.Map {
		if k < key {
			mp[k] = *(ped.Map[k])
			delete(ped.Map, k)
		}
	}

	return
}

// Pedrail start rolling, exce Do() if not nil, and send data to log
// stop rolling when quit closed
func (ped *Pedrail) Start() {
	var fn func(map[int64]int64)
	dur := time.Duration(ped.n*ped.seconds) * time.Second
	ped.infof("start: %s, %v", LogTime("0102T15:04:05"), dur)

	if ped.do == nil {
		fn = func(mp map[int64]int64) {
			return
		}
	} else {
		fn = func(mp map[int64]int64) {
			if err := ped.do(mp); err != nil {
				ped.errorf("do: %s", err.Error())
			}
		}
	}

	go func() {
		for {
			select {
			case <-ped.ticker.C:
				mp := ped.Roll()
				go fn(mp)
			case <-ped.quit:
				ped.infof("stop: %s, %s", LogTime("0102T15:04:05"), ped.DataJSON())
				return
			}
		}
	}()

	return
}

// add number to cuurent clock (key)
func (ped *Pedrail) Add(n int64) (r int64, err error) {
	var ok bool
	var k int64
	var v *int64

	k = ped.currentKey()
	if v, ok = ped.Map[k]; !ok {
		err = fmt.Errorf("key not found: %d", k)
		ped.errorf("key not found: %d", k)
		return 0, err
	}
	r = atomic.AddInt64(v, n)

	return
}

// stop rolling by close quit
func (ped *Pedrail) Stop() {
	ped.ticker.Stop()
	ped.infof("close")
	close(ped.quit)

	return
}

// set Info logger
func (ped *Pedrail) SetInfof(
	fn func(format string, args ...interface{}) (int, error)) (err error) {

	if fn == nil {
		return fmt.Errorf("fn is nil")
	}
	ped.infof = fn

	return
}

// set error logger
func (ped *Pedrail) SetErrorf(
	fn func(format string, args ...interface{}) (int, error)) (err error) {

	if fn == nil {
		return fmt.Errorf("fn is nil")
	}
	ped.errorf = fn

	return
}

func (ped *Pedrail) Keys() (keys []int64) {
	keys = make([]int64, 0, len(ped.Map))
	for k := range ped.Map {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	return
}

// including current key
func (ped *Pedrail) DataJSON() string {
	// str, _ = JsonStr(ped.Data(), false)
	keys := ped.Keys()
	ts := time.Now().Unix()
	slice := make([]string, 0, len(keys))
	for _, k := range keys {
		v := atomic.LoadInt64(ped.Map[k])
		if k <= ts || v > 0 {
			slice = append(slice, fmt.Sprintf(`"%d":%d`, k, v))
		}
	}

	return "{" + strings.Join(slice, ",") + "}"
}

func FloorClock(n int, u string) (t time.Time, err error) {
	if n <= 0 {
		err = fmt.Errorf("invalid parameter n")
		return
	}

	t = time.Now()
	year, month, day := t.Date()
	hour, minute, second := t.Clock()

	switch u {
	case "second", "S":
		if n > 60 {
			err = fmt.Errorf("invaid n for clock unit")
			return
		}
		t = time.Date(year, month, day, hour, minute, second/n*n, 0, t.Location())
	case "minute", "M":
		if n > 60 {
			err = fmt.Errorf("invaid n for clock unit")
			return
		}
		t = time.Date(year, month, day, hour, minute/n*n, 0, 0, t.Location())
	case "hour", "H":
		if n > 24 {
			err = fmt.Errorf("invaid n for clock unit")
			return
		}
		t = time.Date(year, month, day, hour/n*n, 0, 0, 0, t.Location())
	default:
		err = fmt.Errorf("invalid parameter u")
	}

	return
}

func LogTime(clock string) (str string) {
	now := time.Now() // "2006-01-02T15:04:05", "0102T15:04:05"
	return fmt.Sprintf("%s.%06d%s",
		now.Format(clock),
		now.UnixNano()%1e9/1e3,
		now.Format("-0700"),
	)
}
