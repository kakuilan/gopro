//Writer.Buffered()
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
	fmt.Println("写入前写入字节数为:", w.Buffered())

	w.Write(p)
	fmt.Printf("写入%q后,写入的字节数为:%d\n", string(p), w.Buffered())

	w.Flush()
	fmt.Println("执行Flush方法后,写入的字节数为:", w.Buffered())


}
