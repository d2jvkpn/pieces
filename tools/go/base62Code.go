package main

import (
	"fmt"
	"github.com/d2jvkpn/gorover/rover"
	"os"
	"strconv"
)

const USAGE = `Encode/Decode between numeric string and alphanumeric(base62, 0-9A-Za-z)
  $ base62Code  <c|C>  <int64>
  $ base62Code  <d|D>  <string>`

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, USAGE)
		os.Exit(2)
	}

	var num int64
	var err error

	switch os.Args[1] {
	case "c", "C":
		if num, err = strconv.ParseInt(os.Args[2], 10, 64); err == nil {
			fmt.Println(rover.Base62Encode(num))
		}
	case "d", "D":
		if num, err = rover.Base62Decode(os.Args[2]); err == nil {
			fmt.Println(num)
		}
	default:
		err = fmt.Errorf("invalid command %q", os.Args[1])
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
