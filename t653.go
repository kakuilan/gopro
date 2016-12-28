//Reader.ReadBytes(),参数delim用于指定分隔符字节
package main
import (
		"fmt"
		"bytes"
		"bufio"
		)

func main(){
	data := []byte("Hello,world!")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)

	var delim byte = ','
	line,err := r.ReadBytes(delim)
	fmt.Println(string(line), err)


}
