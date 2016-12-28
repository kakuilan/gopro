//bytes.Split
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	s := []byte("Hello,world!")
	sep := []byte(",")
	s2 := bytes.Split(s,sep)

	for _,f := range s2 {
		fmt.Printf("%q\n", f)
	}
	fmt.Println(s2)
}
