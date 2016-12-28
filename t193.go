// net ResolveTCPAddr() 创建一个TCPAddr
package main
import (
    "fmt"
    "net"
)

func main(){
    ip,err := net.ResolveTCPAddr("tcp", "www.baidu.com:80")
    if err != nil {
	fmt.Println(err)
    }
    fmt.Println(ip)

}
