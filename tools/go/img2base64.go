package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var (
		bts []byte
		err error
	)

	if bts, err = ioutil.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	out := base64.StdEncoding.EncodeToString(bts)
	fmt.Printf(out)
}
