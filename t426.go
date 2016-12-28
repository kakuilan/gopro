package main
import "fmt"

func main(){
	//修改字符串也可写为
	s := "hello"
	s = "c" + s[1:] //字符串虽不能更改,但可进行切片操作
	fmt.Printf("%s\n",s)


}
