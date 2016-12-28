package main
import "fmt"
import "net"

func main(){
    conn1, err1 := net.Dial("tcp", "192.168.1.33:6379")
    conn2, err2 := net.Dial("udp", "192.168.1.129:975")
    conn3, err3 := net.Dial("ip4:icmp", "www.baidu.com")
    conn4, err4 := net.Dial("ip4:1", "10.0.0.3")

    fmt.Println(conn1, conn2, conn3, conn4)
    fmt.Println(err1, err2, err3, err4)

str := `
hello word
haha
`
fmt.Println(str)
}
