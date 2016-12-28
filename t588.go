//bytes.ReadBytes()
//参数delim是指定的分隔符
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	buf := []byte("Hello,world!")
	b := bytes.NewBuffer(buf)
	var delim byte=','
	line,err := b.ReadBytes(delim)
	fmt.Println(string(line),err)


}
