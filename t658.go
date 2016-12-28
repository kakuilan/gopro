//Reader.Buffered()
package main
import (
		"fmt"
		"bytes"
		"bufio"
		)

func main(){
	data := []byte("中华人民共和国")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)

	var buf [12]byte
	n,err := r.Read(buf[:]) //第一次执行Read读取4个Unicode字符
	fmt.Println(string(buf[:n]), n, err)

	rn := r.Buffered()
	fmt.Println("Buffered:", rn)

	n,err = r.Read(buf[:]) //第二次读取剩余的3个Unicode字符
	fmt.Println(string(buf[:n]), n, err)

	rn = r.Buffered()
	fmt.Println("Buffered:", rn)

}
