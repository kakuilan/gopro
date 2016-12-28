//Writer.Flush()
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

	w.Write(p)
	fmt.Printf("未执行Flush缓冲区输出%q\n", string(wr.Bytes()))
	w.Flush()
	fmt.Printf("执行Flush后缓冲区输出%q\n", string(wr.Bytes()))


}
