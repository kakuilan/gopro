//Reader.ReadRune()
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	b := []byte("中华人民共和国")
	r := bytes.NewReader(b)

	ch,size,err := r.ReadRune()
	fmt.Println(string(ch), size, err)

}
