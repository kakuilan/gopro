//switch
package main
import "fmt"

func main(){
	i := 10
	switch i {
		case 1:
			fmt.Println("i == 1")
		case 2,3,4:
			fmt.Println("i=2 or i=3 or i=4")
		case 10:
			fmt.Println("i==10")
		default:
			fmt.Println("All i know is that i is an integer")
	
	}


}
