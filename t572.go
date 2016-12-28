//bytes.ToLowerSpecial
package main
import (
		"fmt"
		"bytes"
		"unicode"
		)


func main(){
	s := []byte("Hello,World!")
	fmt.Println(string(bytes.ToLower(s)))
	fmt.Println(string(bytes.ToLowerSpecial(unicode.AzeriCase, s)))

}
