//Reader.ReadRune()
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

	ch,size,err := r.ReadRune()
	fmt.Println(string(ch), size, err)

}
