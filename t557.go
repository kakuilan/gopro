//bytes.Map
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	s := []byte("同学们，上午好！")
	m := func(r rune) rune {
		if r =='上' {
			r = '下'
		}
		return r
	}

	fmt.Println(string(s))
	fmt.Println(string(bytes.Map(m,s)))
}
