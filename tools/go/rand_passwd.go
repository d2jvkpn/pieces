package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456789

const USAGE = `randStr usgage:
$ rand_passwd  <length>  [chars]
  generate randome password

  e.g. $ rand_passwd 32 "@/"
`

func main() {
	var (
		i, length   int
		tmpl        string
		err         error
		bts, result []byte
		rd          *rand.Rand
	)

	inArr := func(strs []string, str string) bool {
		for i := range strs {
			if str == strs[i] {
				return true
			}
		}
		return false
	}

	if len(os.Args) > 1 && inArr([]string{"--help", "-help", "-h"}, os.Args[0]) {
		fmt.Fprintln(os.Stderr, USAGE)
		os.Exit(0)
	}

	if len(os.Args) > 1 {
		length, err = strconv.Atoi(os.Args[1])
		ErrExit(err)
		if length <= 0 {
			ErrExit(fmt.Errorf("invalid length"))
		}
	} else {
		length = 32
	}

	tmpl = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"
	if len(os.Args) > 2 {
		tmpl += os.Args[2]
	}

	rd = rand.New(rand.NewSource(time.Now().UnixNano()))
	bts, result = []byte(tmpl), make([]byte, length)

	for i = 0; i < length; i++ {
		x := bts[int(rd.Int31n(int32(len(bts))))]
		if bytes.Contains(result, []byte{x}) {
			i--
			continue
		}
		result[i] = x
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
