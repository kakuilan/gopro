//defer语句定义函数延迟执行
package main
import "fmt"

func main(){
	f1()
	f2()
	fmt.Printf("\n")
	fmt.Println(f3())

}

//一个延迟执行的函数的变量的值在声明延迟时别赋值
func f1(){
	i := 0
	defer fmt.Println("f1:", i)
	i++
	return
}

//被延迟的函数按照先进后出的顺序执行
func f2(){
	for i:=0;i<4;i++{
		defer fmt.Printf(" f2: %d ", i)
	}
}

//被延迟的匿名函数会读取函数f3的返回值,或对f3的返回值赋值
func f3()(i int) {
	defer func(){
		i++
	}()
	return 1
}

