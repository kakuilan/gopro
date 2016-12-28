//浮点转换整型
package main
import "fmt"

func main(){
	var i,j,k int
	var f1 float32 = 3.14
	var f2 float64 = 123.456
	var f3 float32 = 4.611

	i = int(f1)
	j = int(f2)
	k = int(f3)
	fmt.Println(i,j,k)
}
