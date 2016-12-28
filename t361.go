// 从1.4开始,不再支持多级指针查找方法成员
package main
import "fmt"

type X struct {}

func (*X) test(){
	println("X.test")
}

func main(){
	p := &X{}
	p.test()

	fmt.Println(p)
	//Error:
	(&p).test()

}
