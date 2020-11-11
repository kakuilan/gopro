//unix-socket-服务端
package main

import (
	"log"
	"net"
)

func main() {
	l, err := net.Listen("unix", "/tmp/unix.sock")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		go helloServer(conn)
	}

}

func helloServer(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		log.Println(string(data))

		_, err = c.Write([]byte("hello, i`m server"))
		if err != nil {
			log.Fatal("Write error:", err)
		}
	}
}
