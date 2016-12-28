//bytes.WriteString
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	s := "你好，时间！"
	b := bytes.NewBuffer(nil)
	n,err := b.WriteString(s)

	fmt.Println(string(b.Bytes()),n,err)

}
