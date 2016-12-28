//Reader.Len
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	b := []byte("Golang")
	r := bytes.NewReader(b)
	fmt.Println(r.Len())

}
