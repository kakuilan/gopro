//bytes.ReadByte
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	buf := []byte("Golang")
	b := bytes.NewBuffer(buf)
	c,err := b.ReadByte()
	fmt.Println(string(c),err)


}
