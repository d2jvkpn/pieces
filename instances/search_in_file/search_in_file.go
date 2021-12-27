package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	var (
		debug      bool
		idx        int
		target, fp string
		err        error
		file       *os.File
		start      time.Time
	)

	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatalln("required <match> <file>")
	}

	target, fp = flag.Args()[0], flag.Args()[1]
	start = time.Now()

	if file, err = os.Open(fp); err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if idx, err = SearchText([]byte(target), file, debug); err != nil {
		fmt.Fprintf(os.Stderr, "SearchText: %v\n", err)
		os.Exit(1)
	}
	log.Printf("Elapsed: %s\n", time.Now().Sub(start))

	if idx == -1 {
		fmt.Println("NotFound: -1")
	} else {
		fmt.Printf("Index: %d\n", idx)
	}

}

func SearchText(target []byte, r io.Reader, debug bool) (idx int, err error) {
	var (
		k, t   int // k: number of bytes try to read, t: temporary value
		n, s   int // n: bytes read, s: search position
		data   []byte
		reader *bufio.Reader
	)

	reader = bufio.NewReader(r)
	k = len(target)              // k = 4 or len(target) + 1
	data = make([]byte, 0, 1024) // 10, 24, 32, 1024
	if debug {
		log.Printf(
			">>> target=%q, k=%d, len(target)=%d, cap(data) =%d\n",
			target, len(target), k, cap(data),
		)
	}

	for {
		if t = len(data); t+k > cap(data) {
			idx += len(data)
			t, data = 0, data[:0]
		}
		if debug {
			log.Printf("~~~ t=%d, k=%d\n", t, k)
		}

		if n, err = io.ReadFull(reader, data[t:(t+k)]); err != nil {
			// !! ErrUnexpectedEOF means that EOF was encountered in the middle of reading a
			//    fixed-size block or data structure
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				return -1, nil
			} else { // invalid utf8...
				return -1, err
			}
		}

		data = data[:len(data)+n] // !! extend data
		if debug {
			log.Printf("    n=%d\n", n)
			log.Printf("    read to data[%d:%d]: %q\n", len(data)-n, len(data), string(data))
		}

		if t = len(data) - k - n; t < 0 { // search from the end of data
			t = 0
		}
		if s = bytes.Index(data[t:], target); s >= 0 {
			if debug {
				log.Printf("    found %q: data[%d:%d]\n", target, t, t+len(target))
			}
			return idx + s + t, nil
		}

	}
}
