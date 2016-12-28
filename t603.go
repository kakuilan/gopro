//Reader.UnreadByte()
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	b := []byte("Golang")
	var buf [6]byte
	r := bytes.NewReader(b)

	err := r.UnreadByte() //Buffer还没被读取过,返回一个错误
	fmt.Println(err)

	n,_ := r.Read(buf[:]) //第一次读取
	fmt.Println(string(buf[:n]), n)

	err = r.UnreadByte() //取消第一次读取的最后一个字节'g'
	fmt.Println(err)

	n,_ = r.Read(buf[:]) //第二次读取,Buffer中只剩'g'了
	fmt.Println(string(buf[:n]), n)


}

