//bytes.TrimSpace
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	s := []byte(" hello,world! ")
	fmt.Println(string(bytes.TrimSpace(s)))

}
