// 闭包复制的是原对象指针,这就很容易解释延迟引用现象
package main
import "fmt"


func test() func(){
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)

	return	func(){
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}

func main(){
	f := test()
	f()


}
