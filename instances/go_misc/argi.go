package main

import (
	"bufio"
	"compress/gzip"
	"os"
	"strings"
)

// argument input structure
type Argi struct {
	Scanner *bufio.Scanner
	name    string
	file    *os.File
	reader  *gzip.Reader
}

// provide input command string and return *Argi and error
func NewArgi(name string) (input *Argi, err error) {
	input = new(Argi)
	input.name = name

	if input.name == "-" {
		input.Scanner = bufio.NewScanner(os.Stdin)
		return
	}

	input.file, err = os.Open(input.name)

	if err != nil {
		return
	}

	if strings.HasSuffix(input.name, ".gz") {
		if input.reader, err = gzip.NewReader(input.file); err != nil {
			return
		}
		input.Scanner = bufio.NewScanner(input.reader)
	} else {
		input.Scanner = bufio.NewScanner(input.file)
	}

	return
}

// get input raw name
func (input *Argi) GetName() string {
	return input.name
}

// get input type: stdin, gzip or text
func (input *Argi) GetType() (s string) {
	switch {
	case input.name == "-":
		s = "stdin"
	case strings.HasSuffix(input.name, ".gz"):
		s = "gzip"
	default:
		s = "text"
	}

	return
}

// close reader and file
func (input *Argi) Close() {
	if input.reader != nil {
		input.reader.Close()
	}

	if input.file != nil {
		input.file.Close()
	}
}
