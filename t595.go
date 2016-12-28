//Buffer.UnreadByte()
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	buf := []byte("Golang")
	b := bytes.NewBuffer(buf)

	fmt.Println(string(b.Next(3)))

	b.UnreadByte()
	fmt.Println(string(b.Bytes()))


}
