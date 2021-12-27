package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var (
		debug  bool
		idx    int
		target string
		err    error
		file   *os.File
	)

	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatalln("required <match> <file>")
	}

	target = flag.Args()[0]
	if file, err = os.Open(flag.Args()[1]); err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if idx, err = SearchText(target, file, debug); err != nil {
		fmt.Fprintf(os.Stderr, "SearchText: %v\n", err)
		os.Exit(1)
	} else if idx == -1 {
		fmt.Println("NotFound: -1") // os.Exit(0)
	} else {
		fmt.Printf("Index: %d\n", idx)
	}
}

func SearchText(target string, r io.Reader, debug bool) (idx int, err error) {
	var (
		// k: number of bytes try to read, t: temporary value, n: bytes read, s: search position
		k, t, n, s int
		data       []byte
		reader     *bufio.Reader
	)

	reader = bufio.NewReader(r)
	k = len(target)              // k = 4 or len(target) + 1
	data = make([]byte, 0, 1024) // 10, 24, 32, 1024

	if debug {
		log.Printf(
			"target=%q, k=%d, len(target)=%d, cap(data) =%d\n",
			target, len(target), k, cap(data),
		)
	}

	for {
		if t = len(data); t+k > cap(data) {
			idx += len(data)
			t, data = 0, data[:0]
		}

		if n, err = io.ReadFull(reader, data[t:(t+k)]); err != nil {
			// !! ErrUnexpectedEOF means that EOF was encountered in the middle of reading a fixed-size block or data structure
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				return -1, nil
			} else { // invalid utf8...
				return -1, err
			}
		}
		if debug {
			log.Printf("~~~ t=%d, k=%d, n=%d\n", t, k, n)
		}

		data = data[:len(data)+n]         // !! extend data
		if t = len(data) - k - n; t < 0 { // search from the end of data
			t = 0
		}

		if s = bytes.Index(data[t:], []byte(target)); s > 0 {
			if debug {
				log.Printf("found %q: data[%d:%d]\n", target, t, t+len(target))
			}
			return idx + s + t, nil
		}
		if debug {
			log.Printf("    read data[%d:%d] bytes: %q\n", len(data)-n, len(data), string(data))
		}
	}
}
