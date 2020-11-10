//tcp操作-服务端
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		dat, _, err := reader.ReadLine()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("client:", string(dat))
		_, err = conn.Write([]byte("word\n"))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Start server with: %s", l.Addr())

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}

}
