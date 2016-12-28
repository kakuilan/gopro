package main
import "fmt"

var i = "i"

func main(){
    i := i
    for i := 0;i<2; i++ {
	i := i + 100
	fmt.Println(i)
    }
    fmt.Println("i = ", i)
}
