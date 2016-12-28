package main
import "fmt"

type T struct {
    i int;
    f float64;
    next *T
}

func main(){
    t := new(T)
    fmt.Println(t)
}
