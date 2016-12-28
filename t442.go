//func 变参
package main
import "fmt"

func ff(arg ...int) {
	for _,n := range arg {
		fmt.Println(n)
	}
	fmt.Printf("类型:%T \n", arg)
}

func main(){
	ff(1,2,3,4,5)

}
