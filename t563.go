//字节切片检查前后缀
package main
import (
		"fmt"
		"bytes"
		)


func main(){
	s := []byte("recover")
	prefix := []byte("re")
	suffix := []byte("ed")
	fmt.Println(bytes.HasPrefix(s, prefix))
	fmt.Println(bytes.HasSuffix(s, suffix))
}

