//IP地址类型
package main
import (
	"fmt"
	"net"
	"os"
)

func main(){
	if len(os.Args) !=2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip.addr \n", os.Args[0])
		os.Exit(1)
	}
	addr := os.Args[1]
	ip := net.ParseIP(addr)
	if ip == nil {
		fmt.Println("Invalid address")
	}else{
		fmt.Println("The address is ", ip.String())
	}
	os.Exit(0)

}
