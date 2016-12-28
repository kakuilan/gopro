//Writer.Write()
package main
import (
		"fmt"
		"bytes"
		"bufio"
		)

func main(){
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	p := []byte("Hello,world!")

	n,err := w.Write(p)
	w.Flush()
	fmt.Println(string(wr.Bytes()), n, err)


}
