//Writer.Available()
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
	fmt.Println("写入前未使用的缓冲区为:", w.Available())

	w.Write(p)
	fmt.Printf("写入%q后,未使用的缓冲区为:%d\n", string(p), w.Available())


}
