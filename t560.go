//bytes.Runes
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	s := []byte("中华人民共和国")
	r := bytes.Runes(s)
	fmt.Printf("转换前字符串 %q 长度: %d字节\n", string(s), len(s))
	fmt.Printf("转换后字符串 %q 长度: %d字节\n", string(r), len(r))
	fmt.Println("s ", s)
	fmt.Println("r ", r)
}

