//bytes.WriteRune
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	var r1,r2 rune = '中','国'
	b := bytes.NewBuffer(nil)
	b.WriteRune(r1)
	b.WriteRune(r2)

	fmt.Println(string(b.Bytes()))


}
