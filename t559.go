//bytes.Replace
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	s := []byte("google")
	old := []byte("o")
	new := []byte("oo")
	n := 1
	fmt.Println(string(bytes.Replace(s, old, new, n)))
	fmt.Println(string(bytes.Replace(s, old, new, -1)))

}
