// slice
package main
import "fmt"

func main(){
	data := [...]int{0,1,2,3,4,5,6}
	slice := data[1:4:7] //[low:high:max]

	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))


}



