//查询服务端口号
// go build t708.go
// ./t708 tcp www
package main
import (
	"net"
	"fmt"
	"os"
)

func main(){
	if len(os.Args)!=3 {
		fmt.Fprintf(os.Stderr, "Usage: %s networkType service\n", os.Args[0])
		os.Exit(1)
	}

	networkType := os.Args[1]
	service := os.Args[2]
	port,err := net.LookupPort(networkType, service)
	if err != nil {
		fmt.Println("LookupPort error:", err.Error())
		os.Exit(1)
	}

	fmt.Println("The service port is:", port)
	os.Exit(0)

}
