package main

import (
	"encoding/base64"
	// "fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var (
		bts []byte
		src string
		fp  string
		err error
	)

	src, fp = os.Args[1], os.Args[2]

	if bts, err = ioutil.ReadFile(src); err != nil {
		log.Fatal(err)
	}

	if bts, err = base64.StdEncoding.DecodeString(string(bts)); err != nil {
		log.Fatal(err)
	}

	if err = ioutil.WriteFile(fp, bts, 0644); err != nil {
		log.Fatal(err)
	}
}
