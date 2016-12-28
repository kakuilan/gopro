//函数作为参数
package main
import "fmt"

func main(){
	var a,b int = 3,4
	f := sum
	f6(a,b,f)
}

func f6(a,b int,sum func(int,int) int){
	fmt.Println(sum(a,b))
}

func sum(a,b int) int{
	return a+b
}
