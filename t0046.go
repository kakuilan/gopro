//udp统计在线人数-服务端
package main

import (
	"log"
	"net"
	"time"
)

func main() {
	pc, err := net.ListenPacket("udp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Start server with: %s", pc.LocalAddr())
	defer pc.Close()

	clients := make([]net.Addr, 0)

	go func() {
		for {
			for _, addr := range clients {
				_, err := pc.WriteTo([]byte("pong\n"), addr)
				if err != nil {
					log.Println(err)
				}
			}
			time.Sleep(5 * time.Second)
		}
	}()

	for {
		buf := make([]byte, 256)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			log.Println(err)
			continue
		}

		clients = append(clients, addr)
		log.Println(string(buf[0:n]))
		log.Println(addr.String(), "connecting...", len(clients), "connected")
	}

}
