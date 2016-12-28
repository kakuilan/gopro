//错误处理
package main
import "fmt"

func Test(){
    defer func(){
	if err:= recover();err!=nil {
	    fmt.Println(err)
	}
    }()

    divide(5,0) //程序出错,中断执行
    fmt.Println("end of test.") //该语句不会被执行

}

func divide(a,b int) int {
    return a /b
}

func main(){
    Test()
}
