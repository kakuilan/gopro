// switch 
package main
import "fmt"

func main(){
    x := []int{1,2,3}
	i := 2
	switch i {
		case x[1]:
			println("a")
		case 1,3 :
			println("b")
		default:
			println("c")
	}
	fmt.Println(x)
}
