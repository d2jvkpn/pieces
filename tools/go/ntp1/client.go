package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/d2jvkpn/pieces/pkg/go/misc"
)

func main() {
	var (
		delay  int64
		addr   string
		err    error
		result *misc.NetworkTimeResult
	)

	flag.StringVar(&addr, "addr", "", "request addres")
	flag.Int64Var(&delay, "delay", 10, "delay in millsec")
	flag.Parse()

	if addr == "" {
		log.Fatalf("invalid addr: %q\n", addr)
	}

	if result, err = misc.GetNetworkTime(addr, delay); err != nil {
		log.Println(err)
	}

	fmt.Println(result)
}
