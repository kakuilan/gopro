//TCP Server端设计,服务器使用本地地址,服务端口号为5000,服务器设计工作模式采用循环服务器,对每一个连接请求调用线程handleClient来处理
// go build t710.go
// ./t710

package main
import (
	"fmt"
	"net"
	"os"
)

func main(){
	service := ":5000"
	tcpAddr,err := net.ResolveTCPAddr("tcp",service)
	checkError(err)
	listener,err := net.ListenTCP("tcp",tcpAddr)
	checkError(err)
	for {
		conn,err := listener.Accept()
		if err != nil {
			continue
		}
		handleClient(conn)
		conn.Close()
	}

}


func handleClient(conn net.Conn) {
	var buf [512]byte
	for {
		n,err := conn.Read(buf[0:])
		if err!=nil {
			return
		}
		rAddr := conn.RemoteAddr()
		fmt.Println("Receive from client:",rAddr.String(), string(buf[0:n]))
		_,err2 := conn.Write([]byte("Welcome client!"))
		if err2!=nil {
			return
		}
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

}

