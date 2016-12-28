//defer语句
package main
import "fmt"

func main(){
	defer fmt.Println("The first.")
	fmt.Println("The second.")
}
