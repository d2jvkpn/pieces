package main

import (
	"fmt"
	"log"
	"os"

	"github.com/d2jvkpn/pieces/pkg/go/misc"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "not server address provided")
		os.Exit(1)
	}

	addr := os.Args[1]
	ser, err := misc.NewNetworkTimeServer(addr, 10)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf(">>> Network Time Server listening on: %q\n", addr)
	ser.Run()
}
