//bytes.ReadRune
package main
import (
		"fmt"
		"bytes"
		)


func main(){
	buf := []byte("中华人民共和国")
	b := bytes.NewBuffer(buf)
	r,size,err := b.ReadRune()
	fmt.Println(string(r), size, err)


}
