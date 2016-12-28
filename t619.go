//闭包例子
//1、匿名函数嵌套在函数closures内部
//2、函数closures返回匿名函数
//3、匿名函数引用了自身外部的变量x
//也就是说,当函数closures的内部函数(匿名函数)被函数closures外的一个变量引用时,就创建了一个闭包
package main
import "fmt"

func main(){
	f := closures(10)
	fmt.Println(f(1))
	fmt.Println(f(2))

}

func closures(x int) func(int) int{
	return func(y int) int{
		return x+y
	}
}

