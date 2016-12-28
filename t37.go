package main
import "fmt"

func main(){
    defer func(){
	e := recover()
	if e != nil {
	    fmt.Println("出现问题：", e) //打印出错误
	}
    }()

    one()
    two()


}
