//bytes.TrimLeftFunc
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	s := []byte("hello,world!")
	fmt.Println(string(bytes.TrimLeftFunc(s,f)))
}

func f(a rune) bool{
	if a =='h' || a=='H' {
		return true
	}else{
		return false
	}
}
