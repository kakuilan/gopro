//Reader.ReadSlice
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
	line,err := r.ReadSlice(delim)
	fmt.Println(string(line), err)

	line,err = r.ReadSlice(delim)
	fmt.Println(string(line), err)


	line,err = r.ReadSlice(delim)
	fmt.Println(string(line), err)



}
