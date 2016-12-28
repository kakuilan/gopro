//变参函数的声明和调用
package main
import "fmt"

func main(){
	f1(1,2,3)
	f1(1,2,3,4,5)

}

func f1(args ... int) {
	//变参args实质是个切片
	fmt.Println(args)
}
