//Buffer.UnreadRune()
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	buf := []byte("中华人民共和国")
	b := bytes.NewBuffer(buf)

	b.ReadRune()
	b.ReadRune()
	b.ReadRune()
	b.UnreadRune()
	fmt.Println(string(b.Bytes()))
}
