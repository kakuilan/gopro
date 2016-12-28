//bytes.Trim
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	s := []byte("hello,world!")
	cutset := "hold!"
	fmt.Println(string(bytes.Trim(s,cutset)))

}
