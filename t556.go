//bytes.Count
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	s := []byte("banana")
	sep1 := []byte("ban")
	sep2 := []byte("na")
	sep3 := []byte("a")

	fmt.Println(bytes.Count(s, sep1))
	fmt.Println(bytes.Count(s, sep2))
	fmt.Println(bytes.Count(s, sep3))
	

}

