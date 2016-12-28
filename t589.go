//bytes.ReadFrom
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	buf := []byte("中华人民共和国")
	r := bytes.NewBuffer(buf)
	b := bytes.NewBuffer(nil)

	n,err := b.ReadFrom(r)
	fmt.Println(string(b.Bytes()),n,err)

}
