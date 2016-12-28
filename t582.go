//bytes.WriteByte
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	var c1,c2 byte = 'G','o'
	b := bytes.NewBuffer(nil)
	err := b.WriteByte(c1)
	b.WriteByte(c2)
	fmt.Println(string(b.Bytes()), err)

}
