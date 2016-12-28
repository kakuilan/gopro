//switch的fallthrough会强制执行后面的case
package main
import "fmt"

func main(){
	i := 6
	switch i {
		case 4:
			fmt.Println("i <= 4")
			fallthrough
		case 5:
			fmt.Println("i <=5")
			fallthrough
		case 6:
			fmt.Println("i <=6")
			fallthrough
		case 7:
			fmt.Println("i <=7")
			fallthrough
		case 8:
			fmt.Println("i <=8")
			fallthrough
		default:
			fmt.Println("default case")
	
	}


}
