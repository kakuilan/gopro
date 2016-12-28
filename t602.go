//Reader.Seek()
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	b := []byte("中华人民共和国")
	r := bytes.NewReader(b)

	r.Seek(3,0) //从起始位置偏移3,应该是'华'字
	ch,size,err := r.ReadRune()
	fmt.Println(string(ch), size, err)

	r.Seek(3,1) //从当前位置偏移3，应该是'民'
	ch,size,err = r.ReadRune()
	fmt.Println(string(ch), size, err)

	r.Seek(-6,2) //从尾部偏移-6，应该是'和'
	ch,size,err = r.ReadRune()
	fmt.Println(string(ch), size, err)

}
