// 编写一个简单的echo服务,使其监听于本地的TCP端口8053上.它可以读取一行(以换行符结尾),将该行原样返回然后关闭连接
//运行: go build t270.go
// ./T270
// nc 127.0.0.1 8053
// 然后随便输入

package main
import ("net";"fmt";"bufio")

func main(){
    l,err := net.Listen("tcp", "127.0.0.1:8053")
    if err != nil {
	fmt.Printf("Failure to listen:%s\n", err.Error())
    }

    for {
	if c,err := l.Accept();err ==nil {
	    Echo(c)
	}
    }

}

func Echo(c net.Conn) {
    defer c.Close()
    line,err := bufio.NewReader(c).ReadString('\n')
    if err != nil {
	fmt.Printf("Failure to read: %s\n", err.Error())
	return
    }

    _,err = c.Write([]byte(line))
    if err != nil {
	fmt.Printf("Failure to write: %s\n", err.Error())
	return
    }

}
