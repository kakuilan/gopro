//匿名函数的声明和调用
package main
import "fmt"

func main(){
	//声明并将匿名函数赋值给变量f
	f := func(a,b int) int{
			return a+b
	   }

	//对函数类型变量f进行调用
	sum := f(2,3)
	fmt.Println(sum)

	//声明并直接执行匿名函数
	sum = func(a,b int) int{
		return a+b
	}(5,7)
	fmt.Println(sum)


}
