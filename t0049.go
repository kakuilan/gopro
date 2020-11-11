//unix-socket-客户端
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func reader(r io.Reader) {
	buf := make([]byte, 512)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		log.Println(string(buf[0:n]))
	}
}

func main() {
	c, err := net.Dial("unix", "/tmp/unix.sock")
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()
	go reader(c)

	for {
		_, err := c.Write([]byte("hi, i`m client"))
		if err != nil {
			log.Fatal("Write error:", err)
			break
		}
		time.Sleep(3 * time.Second)
	}

}
