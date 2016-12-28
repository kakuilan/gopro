//Reader.Read()
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

	var buf [128]byte
	n,err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)



}
