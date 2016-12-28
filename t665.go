//Writer.WriteRune()
package main
import (
		"fmt"
		"bytes"
		"bufio"
		)

func main(){
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	var r rune = '中'
	size,err := w.WriteRune(r)
	w.Flush()
	fmt.Println(string(wr.Bytes()), size, err)


}
