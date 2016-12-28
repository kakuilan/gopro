//bytes.Contains函数
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	b := []byte("Golang")
	subslice1 := []byte("go")
	subslice2 := []byte("Go")
	fmt.Println(bytes.Contains(b, subslice1))
	fmt.Println(bytes.Contains(b, subslice2))


}
