//new和make的不同
package main
import "fmt"

func main(){
	var p *[]int = new([]int)
	var v []int = make([]int, 10)

	fmt.Println(p)
	fmt.Println(v)

}
