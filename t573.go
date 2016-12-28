//bytes.ToUpperSpecial
package main
import (
		"fmt"
		"bytes"
		"unicode"
		)

func main(){
	s := []byte("hello,world!")

	fmt.Println(string(bytes.ToUpper(s)))
	fmt.Println(string(bytes.ToUpperSpecial(unicode.AzeriCase, s)))


}
