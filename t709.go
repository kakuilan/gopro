//TCP连接地址
//编译 go build t709.go
//使用 ./t709 tcp www.souhu.com:80
package main
import (
	"fmt"
	"net"
	"os"
)

func main(){
	if len(os.Args)!= 3 {
		fmt.Fprintf(os.Stderr,"Usage: %s networkType addr\n", os.Args[0])
		os.Exit(1)
	}

	networkType := os.Args[1]
	addr := os.Args[2]
	tcpAddr,err := net.ResolveTCPAddr(networkType,addr)
	if err != nil {
		fmt.Println("ResolveTCPAddr error:", err.Error())
		os.Exit(1)
	}

	fmt.Println("The network type is:", tcpAddr.Network())
	fmt.Println("The IP address is ", tcpAddr.String())
	os.Exit(0)

}
