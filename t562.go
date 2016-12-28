//字节切片比较
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	a := []byte("abc")
	b := []byte("Abc")
	s := []byte("GOLANG")
	t := []byte("golang")

	fmt.Println(bytes.Compare(a,b))
	fmt.Println(bytes.Equal(a,b))
	fmt.Println(bytes.EqualFold(s,t))

}
