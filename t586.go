//bytes.Read
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	buf := []byte("Hello,world!")
	b := bytes.NewBuffer(buf)
	var p [8]byte
	n,err := b.Read(p[:])
	fmt.Println(string(p[:n]), n, err)


}
