//获取主机正式名
// go build t706.go
// ./t706 www.163.com


package main
import (
	"fmt"
	"net"
	"os"
)

func main(){
	if len(os.Args)!= 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	cname,err := net.LookupCNAME(name)
	if err != nil {
		fmt.Println("LookupCNAME error:", err.Error())
		os.Exit(1)
	}

	fmt.Println("The host canonical name is:", cname)
	os.Exit(0)

}
