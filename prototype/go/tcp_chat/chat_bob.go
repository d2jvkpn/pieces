package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	var (
		err  error
		addr string
		wg   sync.WaitGroup
		lis  net.Listener
		conn net.Conn
	)

	if len(os.Args) > 1 {
		addr = os.Args[1]
	} else {
		addr = "localhost:2020"
	}

	if lis, err = net.Listen("tcp", addr); err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	fmt.Printf("=== Bob is online, please connect to %s!\n", addr)
	if conn, err = lis.Accept(); err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("=== accepted", conn.RemoteAddr())
	wg.Add(2)
	go func() {
		conn.Write([]byte("Hello, This is Bob!\n"))

		for {
			text, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				log.Println("!!!", err)
				break
			}
			conn.Write([]byte(text))
		}

		wg.Done()
	}()

	go func() {
		for {
			msg, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Println("!!!", err)
				break
			}
			fmt.Println(">>>", strings.TrimSpace(msg))
		}

		wg.Done()
	}()

	wg.Wait()
}
