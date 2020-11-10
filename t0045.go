//tcp操作-客户端
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		_, err := conn.Write([]byte("hello\n"))
		if err != nil {
			log.Fatal(err)
		}
		dat, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("server:", string(dat))
		time.Sleep(5 * time.Second)

	}

}
