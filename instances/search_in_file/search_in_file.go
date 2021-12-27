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
		start      time.Time
	)

	flag.BoolVar(&debug, "debug", false, "run in debug mode")
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatalln("required <match> <file>")
	}

	target, fp = flag.Args()[0], flag.Args()[1]
	start = time.Now()

	if idx, err = SearchInFile(target, fp, debug); err != nil {
		fmt.Fprintf(os.Stderr, "SearchInFile: %v\n", err)
		os.Exit(1)
	}
	log.Printf("Elapsed: %s\n", time.Now().Sub(start))

	if idx == -1 {
		fmt.Println("NotFound: -1")
	} else {
		fmt.Printf("Index: %d\n", idx)
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

	return SearchText([]byte(target), file, 32, debug)
}

func SearchText(bts []byte, r io.Reader, bufsize int, debug bool) (idx int, err error) {
	var (
		// k: number of bytes try to read, t: temporary value
		// n: bytes read, s: search position
		k, t, n, s   int
		buffer, tail []byte
		reader       *bufio.Reader
	)

	reader = bufio.NewReader(r)
	k, buffer = 2*len(bts), make([]byte, 0, bufsize)
	if k < bufsize {
		k = bufsize
	}

	if debug {
		log.Printf(
			">>> target=%q, k=%d, len(target)=%d, cap(data) =%d\n",
			bts, len(bts), k, cap(buffer),
		)
	}

	for {
		if t = len(buffer); t+k > cap(buffer) {
			tail = buffer[t-k : t] // !! left shift
			idx += t - k
			// buffer = make([]byte, 0, bufsize)
			// buffer = append(buffer, tail...)
			buffer = append(buffer[:0], tail...) // avoid allocation
		}
		if debug {
			log.Printf("~~~ read to buffer: t=%d, k=%d\n", t, k)
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

			if debug {
				log.Printf("    n=%d, length=%d\n", n, len(buffer))
			}
			return -1, nil
		}

		buffer = buffer[:len(buffer)+n] // !! extend buffer
		if debug {
			log.Printf("    n=%d, length=%d\n", n, len(buffer))
			log.Printf("    buffer=%q\n", string(buffer[:len(buffer)]))
		}

		if t = len(buffer) - k - n; t < 0 { // search from the end of buffer
			t = 0
		}
		if s = bytes.Index(buffer[t:], bts); s >= 0 {
			idx = idx + s + t
			if debug {
				log.Printf("<<< found %q: range=[%d:%d]\n", bts, idx, idx+n)
			}
			return idx, nil
		}
	}
}
