//bytes.FieldsFunc
package main
import (
		"fmt"
		"bytes"
		)


func main(){
	s := []byte("小明和爸爸和妈妈")
	s2 := bytes.FieldsFunc(s,f)

	for _,f := range s2 {
		fmt.Printf("%q\n", f)
	}
	fmt.Println(s2)
}

func f(a rune) bool {
	if a =='和' {
		return true
	}else{
		return false
	}
}
