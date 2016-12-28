//Buffer.Bytes(),Buffer.Len()
package main
import (
		"fmt"
		"bytes"
		)


func main(){
	buf := []byte("Golang is a beautiful language, I like it!")
	b := bytes.NewBuffer(buf)
	var delim byte = ','
	fmt.Println("未执行读操作前Buffer中数据和字节数")
	fmt.Println(string(b.Bytes()), b.Len())

	b.ReadString(delim)
	fmt.Println("执行读操作后Buffer中数据和字节数")
	fmt.Println(string(b.Bytes()), b.Len())


}


