package main
import "fmt"

func main(){
    add := func(base int) func(int) int{
	return func(n int) int{
	    return base + n
	}
    }

    add5 := add(5)
    fmt.Println(add5(10))
}
