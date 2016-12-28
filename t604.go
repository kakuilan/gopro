//Reader.UnreadRune()
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	b := []byte("中华人民共和国")
	var buf [3]byte
	r := bytes.NewReader(b)

	r.Read(buf[:]) //用Read读取第一个Unicode字符'中'
	err := r.UnreadRune()
	fmt.Println(err)

	ch,_,_ := r.ReadRune() //用ReadRune读取第二个Unicode字符'华'
	fmt.Println(string(ch))

	err = r.UnreadRune() //取消上次读取的Unicode字符
	fmt.Println(err)

	ch,_,_ = r.ReadRune() //用ReadRune重新读取第二个Unicode字符
	fmt.Println(string(ch))


}
