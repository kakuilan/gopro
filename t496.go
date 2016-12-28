//整型到浮点
package main
import "fmt"

func main(){
	var i int = 100
	var f1 float32
	var f2 float64
	f1 = float32(i)
	f2 = float64(i)
	fmt.Println(f1, f2)
	fmt.Printf("%f , %f\n", f1, f2)

}
