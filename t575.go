//bytes.TrimFunc
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	s := []byte("hello,world!")
	fmt.Println(string(bytes.TrimFunc(s,f)))
}

func f(a rune) bool {
	if a == 'h'||a=='!' {
		return true
	}else{
		return false
	}
}
