package main
import "fmt"

func example(x int) int{
    if x ==0 {
	return 5
    } else {
	return x
    }
}

func main(){
    var a int
    a = example(9)
    fmt.Println(a)

}
