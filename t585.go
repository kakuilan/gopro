//bytes.WriteTo
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	buf := []byte("Golang")
	b := bytes.NewBuffer(buf)
	w := bytes.NewBuffer(nil)
	n,err := b.WriteTo(w)
	fmt.Println(string(w.Bytes()),n,err)


}
