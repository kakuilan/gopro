//Reader.ReadAt()
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	b := []byte("Golang")
	var buf [3]byte

	r := bytes.NewReader(b)
	n,err := r.ReadAt(buf[:], 2)
	fmt.Println(string(buf[:n]), n, err)

	n,err = r.ReadAt(buf[:], 3)
	fmt.Println(string(buf[:n]), n, err)

	n,err = r.ReadAt(buf[:], 4)
	fmt.Println(string(buf[:n]), n, err)



}
