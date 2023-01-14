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
		err     error
		addr    string
		wg      sync.WaitGroup
		tcpaddr *net.TCPAddr
		conn    *net.TCPConn
	)

	if len(os.Args) > 1 {
		addr = os.Args[1]
	} else {
		addr = "localhost:2020"
	}

	if tcpaddr, err = net.ResolveTCPAddr("tcp", addr); err != nil {
		log.Fatal(err)
	}

	if conn, err = net.DialTCP("tcp", nil, tcpaddr); err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("=== Alice is online, connected to", conn.RemoteAddr())
	if _, err = conn.Write([]byte("Hello Bob, I'm Alice!\n")); err != nil {
		log.Fatalln(err)
	}

	wg.Add(2)
	go func() { // alice input
		for {
			text, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				log.Println("!!!", err)
				break
			}
			fmt.Println("=>>", text)
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
			fmt.Println("<<=", strings.TrimSpace(msg))
		}

		wg.Done()
	}()

	wg.Wait()
}
