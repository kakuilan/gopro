//Reader.ReadByte()
package main
import (
		"fmt"
		"bytes"
		"bufio"
		)

func main(){
	data := []byte("Golang")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)

	c,err := r.ReadByte()
	fmt.Println(string(c), err)


}
