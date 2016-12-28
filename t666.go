//Writer.WriteString()
package main
import (
		"fmt"
		"bytes"
		"bufio"
		)

func main(){
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	s := "Hello,world!"
	n,err := w.WriteString(s)
	w.Flush()
	fmt.Println(string(wr.Bytes()), n, err)

}
