//bytes.Join
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	s := [][]byte{
		[]byte("你好"),
		[]byte("世界!"),
	   }
	sep := []byte(",")
	fmt.Println(string(bytes.Join(s, sep)))

}
