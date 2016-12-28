//Reader.Peek()
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

	bl,err := r.Peek(9)
	fmt.Println(string(bl), err)

	bl,err = r.Peek(18)
	fmt.Println(string(bl), err)

	bl,err = r.Peek(27)
	fmt.Println(string(bl), err)


}
