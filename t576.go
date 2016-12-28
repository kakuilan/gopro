//bytes.TrimLeft
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	s := []byte("hello,world!")
	cutset := "hold"
	fmt.Println(string(bytes.TrimLeft(s, cutset)))
}
