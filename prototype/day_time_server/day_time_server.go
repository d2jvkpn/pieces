/* DaytimeServer
 */
package main

import (
	"log"
	"net"
	"os"
	"time"
)

func main() {
	var (
		addr    string
		err     error
		tcpAddr *net.TCPAddr
		lis     *net.TCPListener
	)

	if len(os.Args) > 1 {
		addr = os.Args[1]
	} else {
		addr = ":2020"
	}
	if tcpAddr, err = net.ResolveTCPAddr("tcp4", addr); err != nil {
		log.Fatalln(err)
	}

	if lis, err = net.ListenTCP("tcp", tcpAddr); err != nil {
		log.Fatalln(err)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	log.Printf("RemoteAddr: %s\n", conn.RemoteAddr())

	conn.Write([]byte(time.Now().Format("2006-01-02T15:04:05.000Z07:00\n")))
	conn.Close()
}
