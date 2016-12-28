// net InterfaceAddrs() 返回本机网络地址列表
package main
import (
    "fmt"
    "net"
)

func main(){
    addr,err := net.InterfaceAddrs()
    if err != nil {
	fmt.Println(err)
    }
    fmt.Println(addr)


}
