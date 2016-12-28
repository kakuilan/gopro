//Writer.WriteByte()
package main
import (
		"fmt"
		"bytes"
		"bufio"
		)

func main(){
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)

	var c byte = 'G'
	err := w.WriteByte(c)
	w.Flush()
	fmt.Println(string(wr.Bytes()), err)


}
