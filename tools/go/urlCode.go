package main

import (
	"fmt"
	"net/url"
	"os"
)

const USAGE = `URL Escape/Unescape, usage:
  $ urlCode  <e|E>  <string>
  $ urlCode  <u|U>  <url>
`

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, USAGE)
		os.Exit(2)
	}

	var (
		cmd, str, result string
		err              error
	)

	cmd, str = os.Args[1], os.Args[2]

	switch cmd {
	case "e", "E":
		fmt.Println(url.QueryEscape(str))
	case "u", "U":
		if result, err = url.QueryUnescape(str); err == nil {
			fmt.Println(result)
		}
	default:
		err = fmt.Errorf("invalid command %q", cmd)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
