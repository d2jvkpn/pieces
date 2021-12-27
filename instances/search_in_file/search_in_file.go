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
		// k: number of bytes try to read, t: temporary value
		// n: bytes read, s: search position
		k, t, n, s int
		buffer     []byte
		reader     *bufio.Reader
	)

	reader = bufio.NewReader(r)
	k = len(target)                // k = 4 or len(target) + 1
	buffer = make([]byte, 0, 1024) // 10, 24, 32, 1024
	if debug {
		log.Printf(
			">>> target=%q, k=%d, len(target)=%d, cap(data) =%d\n",
			target, len(target), k, cap(buffer),
		)
	}

	for {
		if t = len(buffer); t+k > cap(buffer) {
			idx += len(buffer)
			t, buffer = 0, buffer[:0]
		}
		if debug {
			log.Printf("~~~ read to buffer: t=%d, k=%d\n", t, k)
		}

		if n, err = io.ReadFull(reader, buffer[t:(t+k)]); err != nil {
			// !! ErrUnexpectedEOF means that EOF was encountered in the middle of reading a
			//    fixed-size block or data structure
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				log.Printf("    n=%d, length=%d\n", n, len(buffer))
				return -1, nil
			} else { // invalid utf8...
				return -1, err
			}
		}

		buffer = buffer[:len(buffer)+n] // !! extend buffer
		if debug {
			log.Printf("    n=%d, length=%d\n", n, len(buffer))
			log.Printf("    buffer: %q\n", string(buffer))
		}

		if t = len(buffer) - k - n; t < 0 { // search from the end of buffer
			t = 0
		}
		if s = bytes.Index(buffer[t:], target); s >= 0 {
			idx = idx + s + t
			if debug {
				log.Printf("<<< found %q: range=[%d:%d]\n", target, idx, idx+n)
			}
			return idx, nil
		}
	}
}
