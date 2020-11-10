//使用scoket访问百度
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	//尝试与百度建立tcp连接
	conn, err := net.Dial("tcp", "baidu.com:443")
	if err != nil {
		log.Fatal(err)
	}

	//退出关闭连接
	defer conn.Close()

	//通过连接发送GET请求,访问首页
	_, err = fmt.Fprintf(conn, "Get / HTTP/1.0\r\n\r\n")
	if err != nil {
		log.Fatal(err)
	}

	dat, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(dat))

}
