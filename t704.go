//获取主机地址信息
// 编译 build t704.go
//使用 ./t704 163.com

package main
import (
	"fmt"
	"net"
	"os"
)

func main(){
	if len(os.Args)!=2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	addrs,err := net.LookupHost(name)
	if err !=nil {
		fmt.Println("LookupHost error:", err.Error())
		os.Exit(1)
	}

	for _,s := range addrs {
		fmt.Println(s)
	}
	os.Exit(0)


}
