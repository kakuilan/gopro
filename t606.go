//slice应用.统计文件tmp中有多少个单词,tmp内容为
// Hello,world! Golang is a beautiful language. I like it!
//解决
//1、定义字节切片buf
//2、使用os包中的File对象的Read方法,将文件的内容读入到字节切片buf中
//3、使用bytes包中的FieldsFunc()函数对文件内容进行断句,比如将标点符号"."、","、"!"作为断句的标识符
//4、对每个断句再使用bytes包中的Fields()函数进行分割,这样独立的单词就被分离出来了,每分离出一个单词,单词计数num+1
package main
import (
		"fmt"
		"bytes"
		"os"
		)

func main(){
	var num int
	buf := make([]byte, 1024)
	
	f,_ := os.Open("tmp606")
	defer f.Close()

	n,_ := f.Read(buf)
	fmt.Println(string(buf[:n]), n)

	for _,sentence := range bytes.FieldsFunc(buf[:n], f1) {
		for _,words := range bytes.Fields(sentence) {
			num++
			fmt.Printf(" %q", words)
		}
	}
	fmt.Println()
	fmt.Println("文件", f.Name(), "里总共有", num, "个单词.")

}

func f1(a rune) bool {
	if a ==',' || a=='.' || a=='!' {
		return true
	}else{
		return false
	}

}
