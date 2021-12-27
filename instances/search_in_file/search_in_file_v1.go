package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var (
		n, t, k int
		data    []byte
		target  string
		err     error
		file    *os.File
		reader  *bufio.Reader
	)

	if len(os.Args) < 3 {
		log.Fatalln("required <match> <file>")
	}

	target = os.Args[1]
	if file, err = os.Open(os.Args[2]); err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	k = len(target) + 1          // k = 4, one for \n
	data = make([]byte, 0, 1024) // 10, 24, 32, 1024
	log.Printf(
		"target=%q, len(target) = %d, k=%d, cap(data) =%d\n",
		target, len(target), k, cap(data),
	)

	reader = bufio.NewReader(file)

	for {
		// n, err = io.ReadFull(reader, buffer[:cap(buffer)])
		if t = len(data); t+k > cap(data) {
			log.Fatal("data is full")
		}

		if n, err = io.ReadFull(reader, data[t:(t+k)]); err != nil {
			switch err {
			case io.EOF:
				log.Println("io.EOF")
				return
			case io.ErrUnexpectedEOF:
				break // continue loop
			default:
				log.Fatal(err)
			}
		}

		data = data[:len(data)+n]         // !! extend data
		if t = len(data) - k - n; t < 0 { // search from the end of data
			t = 0
		}

		if t = bytes.Index(data[t:], []byte(target)); t > 0 {
			log.Printf("found target %q: data[%d:%d]\n", target, t, t+len(target))
			break
		}

		fmt.Printf(
			"read data[%d:%d] bytes: %q\n",
			len(data)-n, len(data), string(data),
		)
	}
}

// https://play.golang.org/p/iYOY-z7hkoz
func demo01() {
	testname := func(stride int) string {
		f, err := ioutil.TempFile("", "test.stride.")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		_, err = f.Write(make([]byte, 2*stride+stride/2))
		if err != nil {
			panic(err)
		}
		return f.Name()
	}

	stride := 1024
	filename := testname(stride)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	buf := make([]byte, 0, stride)
	for {
		n, err := io.ReadFull(r, buf[:cap(buf)])
		buf = buf[:n]
		if err != nil {
			if err == io.EOF {
				break
			}
			if err != io.ErrUnexpectedEOF {
				fmt.Fprintln(os.Stderr, err)
				break
			}
		}

		fmt.Println("read n bytes...", n)
		// process buf
	}
}
