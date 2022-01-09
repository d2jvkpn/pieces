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
		index      int
		target, fp string
		err        error
		startTime  time.Time
	)

	flag.BoolVar(&debug, "debug", false, "run in debug mode")
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatalln("required <match> <file>")
	}

	target, fp = flag.Args()[0], flag.Args()[1]
	startTime = time.Now()

	if index, err = SearchInFile(target, fp, debug); err != nil {
		fmt.Fprintf(os.Stderr, "SearchInFile: %v\n", err)
		os.Exit(1)
	}
	log.Printf("Elapsed: %s\n", time.Now().Sub(startTime))

	if index == -1 {
		fmt.Println("NotFound: -1")
	} else {
		fmt.Printf("Index: %d\n", index)
	}
}

func SearchInFile(target string, fp string, debug bool) (int, error) {
	var (
		err  error
		file *os.File
	)

	if file, err = os.Open(fp); err != nil {
		return -1, err
	}
	defer file.Close()

	return SearchText([]byte(target), file, debug)
}

func SearchText(bts []byte, r io.Reader, debug bool) (index int, err error) {
	var (
		ok bool
		// k: number of bytes try to read, t: temporary value
		// n: bytes read, s: search position
		k, t, n, s   int
		buffer, tail []byte
		reader       *bufio.Reader
	)

	reader = bufio.NewReader(r)
	ok, k = true, len(bts)
	buffer = make([]byte, 0, 8*k)

	if debug {
		log.Printf(">>> target=%q, k=%d, cap(buffer)=%d\n", bts, k, cap(buffer))
	}

	for ok {
		if t = len(buffer); t+k > cap(buffer) {
			tail = buffer[t-k : t] // !! left shift
			index += t - k
			// buffer = make([]byte, 0, bufsize)
			// buffer = append(buffer, tail...)
			buffer = append(buffer[:0], tail...) // avoid allocation
		}

		if debug {
			log.Printf("~~~ read to buffer: [%d:%d], index=%d\n", t, t+k, index)
		}

		if t = len(buffer) + k; t > cap(buffer) {
			t = cap(buffer)
		}
		if n, err = io.ReadFull(reader, buffer[len(buffer):t]); err != nil {
			// !! ErrUnexpectedEOF means that EOF was encountered in the middle of reading a
			//    fixed-size block or data structure
			if err != io.EOF && err != io.ErrUnexpectedEOF {
				return -1, err
			}

			ok = false // don't continue next loop, but execute flowing codes
		}
		buffer = buffer[:len(buffer)+n] // !! extend buffer

		if debug {
			log.Printf("    save to buffer: [%d:%d]\n", len(buffer)-n, len(buffer))
			log.Printf("    buffer=%q\n", string(buffer[:len(buffer)]))
		}

		if t = len(buffer) - k - n; t < 0 { // search from the end of buffer
			t = 0
		}
		if s = bytes.Index(buffer[t:], bts); s >= 0 {
			index += (s + t)

			if debug {
				log.Printf("<<< found %q: range=[%d:%d]\n", bts, index, index+n)
			}

			return index, nil
		}
	}

	return -1, nil
}
