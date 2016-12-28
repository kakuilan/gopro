//bytes.ReadString()
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	buf := []byte("Golang is a beautiful language,I like it!")
	b := bytes.NewBuffer(buf)
	var delim byte = ','
	line,err := b.ReadString(delim)
	fmt.Println(line, err)


}
