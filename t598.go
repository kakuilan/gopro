//Reader.Read()
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	b := []byte("Golang")
	var buf [3]byte
	r := bytes.NewReader(b)
	n,err := r.Read(buf[:])

	fmt.Println(string(buf[:n]),n,err)

	n,err = r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)


	n,err = r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)

}
