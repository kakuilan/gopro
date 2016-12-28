//buffer.Write
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	p := []byte("Hello,world!")
	b := bytes.NewBuffer(nil)
	n,err := b.Write(p)
	fmt.Println(string(b.Bytes()), n, err)


}
