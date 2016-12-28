//bytes.Fields
package main
import (
		"fmt"
		"bytes"
		)


func main(){
	s := []byte("I'm a student.")
	s2 := bytes.Fields(s)
	for _,f := range s2 {
		fmt.Printf("%q\n", f)
	}
	fmt.Println(s2)

}
