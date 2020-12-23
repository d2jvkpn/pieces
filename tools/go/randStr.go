package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456789

const USAGE = `randStr usgage:
$ randStr  [template]  [length]

  template elements example: "[0-9]", "[a-z]", ".-_@"
  length, output string length, example: 5, 10

  e.g. $ randStr "[0-9][A-Z].-" 10
`

func main() {
	var (
		i, length   int
		tmpl        string
		err         error
		bts, result []byte
		rd          *rand.Rand
	)

	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, USAGE)
		os.Exit(2)
	}

	if tmpl = os.Args[1]; len(tmpl) == 0 {
		ErrExit(fmt.Errorf("invalid binary tmpl"))
	}

	fmt.Fprintf(os.Stderr, ">> Input template: %q\n", tmpl)

	tmpl = strings.Replace(tmpl, "[a-z]", "abcdefghijklmnopqrstuvwxyz", -1)
	tmpl = strings.Replace(tmpl, "[A-Z]", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", -1)
	tmpl = strings.Replace(tmpl, "[0-9]", "0123456789", -1)

	length, err = strconv.Atoi(os.Args[2])
	ErrExit(err)
	if length <= 0 {
		ErrExit(fmt.Errorf("invalid length"))
	}

	rd = rand.New(rand.NewSource(time.Now().UnixNano()))
	bts, result = []byte(tmpl), make([]byte, length)

	for i = 0; i < length; i++ {
		result[i] = bts[int(rd.Int31n(int32(len(bts))))]
	}

	fmt.Println(string(result))
}

func ErrExit(err error) {
	if err == nil {
		return
	}

	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}
