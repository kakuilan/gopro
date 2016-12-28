//bytes.TrimRightFunc
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	s := []byte("hello,world!")
	fmt.Println(string(bytes.TrimRightFunc(s,f)))

}

func f(a rune) bool{
	if a =='!' || a=='.' {
		return true
	}else{
		return false
	}
}
