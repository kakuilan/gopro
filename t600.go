//Reader.ReadByte()
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	b := []byte("Golang")
	r := bytes.NewReader(b)

	c,err := r.ReadByte()
	fmt.Println(string(c), err)


}
