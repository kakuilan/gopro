//捕获函数recover只有在延迟调用内直接调用才会终止错误,否则总是返回nil.任何未捕获的错误都会沿调用堆栈向外传递
package main
import "fmt"

func test(){
	defer recover()	//无效
	defer fmt.Println(11,recover())	//无效

	defer func(){
		func(){
			println("defer inner")
			recover()	//无效
		}()
		//err := recover() //有效
		//fmt.Println(222, err)
	}()

	panic("test panic") //最终,错误向外传递
}

func main(){
	test()
}
