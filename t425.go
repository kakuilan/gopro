// go中字符串是不可变的
// 但如果真要修改怎么办呢
package main
import "fmt"

func main(){
	var s string = "hello"
	//错误
	//s[0] = `c`

	//下面的可以实现
	c := []byte(s) //将字符串s转换为[]byte类型
	c[0] = 'c'
	s2 := string(c) //再转换回string类型
	fmt.Println(s2)

	//GO可以使用+操作符来连接2个字符串
	m := ", world"
	a := s+m
	fmt.Println(a)



}
