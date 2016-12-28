package main
import "fmt"

func g(oo int) int{
    oo--
    return oo
}

func main(){
    o := 42
    c := g(o)
    c++
    fmt.Println(c, o)
}
