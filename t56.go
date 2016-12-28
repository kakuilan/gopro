package main
import "fmt"

func main(){
    type Point [2]float32
    type Line [2]Point

    a := Point{1, 2}
    b := Point{3, 4}
    l := Line{a,b}
    fmt.Printf("%g, %g, %v, %v", a[0], l[1][1], b, l)
}
