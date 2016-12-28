//bytes.TrimRight
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	s := []byte("hello,world!")
	cutset := "hold!"
	fmt.Println(string(bytes.TrimRight(s,cutset)))


}
