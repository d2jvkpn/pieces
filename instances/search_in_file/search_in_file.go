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

	if idx, err = SearchText(file, target, debug); err != nil {
		fmt.Fprintf(os.Stderr, "SearchText: %v\n", err)
		os.Exit(1)
	} else if idx == -1 {
		fmt.Println("NotFound: -1") // os.Exit(0)
	} else {
		fmt.Printf("Index: %d\n", idx)
	}
}

func SearchText(r io.Reader, target string, debug bool) (idx int, err error) {
	var (
		// k: number of bytes try to read, t: temporary value, n: bytes read, s: search position
		k, t, n, s int
		data       []byte
		reader     *bufio.Reader
	)

	reader = bufio.NewReader(r)
	k = len(target)              // k = 4, one for \n
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
			if err == io.EOF {
				return -1, nil
			} else { // io.ErrUnexpectedEOF, invalid utf8...
				return -1, err
			}
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
			log.Printf("read data[%d:%d] bytes: %q\n", len(data)-n, len(data), string(data))
		}
	}
}
