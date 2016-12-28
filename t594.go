//Buffer.Truncate()
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	buf := []byte("Hello,world!")
	b := bytes.NewBuffer(buf)
	fmt.Println("未执行Reset前Buffer中内容:", string(b.Bytes()))

	b.Truncate(5)
	fmt.Println("执行Truncate后Buffer中内容:", string(b.Bytes()))

	b.Reset()
	fmt.Println("执行Reset后Buffer中内容:", string(b.Bytes()))


}
