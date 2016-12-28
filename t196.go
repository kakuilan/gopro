// udp test服务端
package main
import (
    "fmt"
    "net"
    "time"
)

func main(){
    //在7070端口监听
    addr,err := net.ResolveUDPAddr("udp", ":7070")
    if err != nil {
	fmt.Println(err)
	return
    }

    conn,err := net.ListenUDP("udp", addr)
    if err != nil {
	fmt.Println(err)
	return
    }

    fmt.Println("UDP服务器启动...")
    for {
	var buf [1024]byte
	n,addr,err := conn.ReadFromUDP(buf[0:])
	if err != nil {
	    fmt.Println(err)
	    return
	}
	go HandleClient(conn, buf[0:n], addr)
    }
}

func HandleClient(conn *net.UDPConn, data []byte, addr *net.UDPAddr) {
    fmt.Println("收到数据:"+string(data))
    conn.WriteToUDP([]byte("OK,数据已收到,接受时间 "+ time.Now().String()), addr)
}
