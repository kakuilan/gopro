//Reader.ReadString()
package main
import (
		"fmt"
		"bytes"
		"bufio"
		)

func main(){
	data := []byte("apple,banana,pear")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)

	var delim byte = ','
	line,err := r.ReadString(delim)
	fmt.Println(line,err)



}
