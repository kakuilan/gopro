//bytes.ToTitleSpecial
package main
import (
		"fmt"
		"bytes"
		"unicode"
		)

func main(){
	s := []byte("hello,world!")

	fmt.Println(string(bytes.Title(s)))
	fmt.Println(string(bytes.ToTitle(s)))
	fmt.Println(string(bytes.ToTitleSpecial(unicode.AzeriCase, s)))


}
