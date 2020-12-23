package main

import (
	"fmt"
	"github.com/d2jvkpn/gorover/rover"
	"os"
	"regexp"
	"strings"
	"time"
)

const USAGE = `Convert between time and timestamp (integer):
  $ tsConvert  <s2t|m2t|u2t|n2t>  <int64>
  $ tsConvert  <t2s|t2m|t2m|t2n>  <string>

  s: Second, m: Millisecond, u: Microsecond, n: Nanosecond
  t string format: 2006-01-02T15:04:05-0700

  e.g.
    $ tsConvert  t2s  ""
    $ tsConvert  t2s  "2020-01-07T23:00:00+0800"
    $ tsConvert  t2s  "2020-01-07 23:00:00 +0800"
    $ tsConvert  t2s  "2020-01-07 23:00:00"
    $ tsConvert  s2t  1578409200

project: github.com/d2jvkpn/gorover`

const TIMEFORMAT = "2006-01-02T15:04:05-0700"

func main() {
	var (
		cmd, str string
		ok       bool
		slice    []string
		now      time.Time
		err      error
		result   string
	)

	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, USAGE)
		os.Exit(2)
	}

	defer func() {
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}()

	cmd, str, now = os.Args[1], os.Args[2], time.Now()
	if strings.HasPrefix(cmd, "t2") && str == "" {
		str = now.Format(TIMEFORMAT)
	}

	slice = strings.Fields(str)
	str = strings.Join(slice, " ")
	if ok, _ = regexp.Match("[0-9]{2}:[0-9]{2}:[0-9]{2}$", []byte(str)); ok {
		str += now.Format("-0700")
		// println(">>> 1", str)
	}

	switch len(slice) {
	case 1:
	case 2:
		str = strings.Replace(str, " ", "T", 1) // concate date and clock
	case 3:
		str = strings.Replace(str, " ", "T", 1)
		str = strings.Replace(str, " ", "", 1) // concate clock and timezone
	default:
		// println(">>> 2", len(slice))
		if err = fmt.Errorf("cann't parse: %q", str); err != nil {
			return
		}
	}

	// fmt.Println(">>> 3", cmd, str)
	if result, err = rover.TimeConvert(cmd, str); err != nil {
		return
	}

	fmt.Println(result)
}
