//函数多返回值
package main
import "fmt"

func main(){
	sum,div := f2(3,6)
	fmt.Println(sum,div)

	//返回值忽略,使用'_'
	rst,_ := f2(5,7)
	fmt.Println(rst)
}


func f2(a,b int)(int,float32) {
	return a*b,float32(a)/float32(b)
}
