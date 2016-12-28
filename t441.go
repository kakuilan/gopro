//func 多个返回值
//如果你的函数是导出的(首字母大写),官方建议:最好命名返回值,因为不命名返回值,虽然代码可能更简介,但是会造成生成的文档可读性差.

package main
import "fmt"

//func SumAndProduct(A,B int) (int,int) {
//	return A+B, A*B
//}

//命名返回值
func SumAndProduct(A,B int) (add int, Multiplied int) {
	add = A+B
	Multiplied = A * B
	return
}

func main(){
	x := 3
	y := 4

	xPLUSy,xTIMESy := SumAndProduct(x,y)

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)


}
