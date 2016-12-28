//TCP Client端设计
// go build t711.go
// ./t711 127.0.0.1:5000

package main
import (
	"fmt"
	"net"
	"os"
)

func main(){
	var buf [512]byte
	if len(os.Args)!=2{
		fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
	}

	service := os.Args[1]
	tcpAddr,err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	conn,err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	rAddr := conn.RemoteAddr()
	n,err := conn.Write([]byte("Hello server!"))
	checkError(err)
	n,err = conn.Read(buf[0:])
	checkError(err)
	fmt.Println("Reply form server:", rAddr.String(), string(buf[0:n]))
	conn.Close()
	os.Exit(0)
	
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr,"Fatal error:%s", err.Error())
		os.Exit(1)
	}
}

