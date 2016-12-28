//指针传递
package main
import "fmt"

func main(){
	var b int = 1
	f2(&b)
	fmt.Println(b)
}

func f2(a *int) {
	*a = 2
	fmt.Printf("%d\n", *a)
}
