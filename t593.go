//Buffer.Next()
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	buf := []byte("Golang")
	b := bytes.NewBuffer(buf)
	fmt.Println(string(b.Next(4)))
	fmt.Println(string(b.Next(4)))
	



}
