package main

import (
	"fmt"
	"os"
	"strconv"
)

const USAGE = `Encode/Decode between numeric string and alphanumeric(base36, 0-9A-Z)
  $ base36Code  <c|C>  <int64>
  $ base36Code  <d|D>  <string>`

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, USAGE)
		os.Exit(2)
	}

	var (
		err error
		num int64
		str string
	)

	str = os.Args[2]

	switch os.Args[1] {
	case "c", "C":
		if num, err = strconv.ParseInt(str, 10, 64); err != nil {
			break
		}

		fmt.Println(strconv.FormatInt(num, 36))
	case "d", "D":
		if num, err = strconv.ParseInt(str, 36, 64); err == nil {
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
