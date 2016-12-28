package main
import (
    "fmt"
    "sort"
)

type Vec struct {x, y float64}

func (v Vec) lenlen() float64 {
    return v.x * v.x + v.y * v.y
}

type VecSlice []Vec
func (v VecSlice) Len() int {
    return len(v)
}
func (v VecSlice) Less(i,j int) bool {
    return v[i].lenlen() < v[j].lenlen()
}
func (v VecSlice) Swap(i,j int) {
    v[i],v[j] = v[j],v[i]
}

func main(){
    v0 := Vec{0,0}
    v1 := Vec{1,0}
    v2 := Vec{1,1}
    vs := VecSlice{v1, v2, v0}
    fmt.Println(vs)
    sort.Sort(vs)
    fmt.Println(vs)


}

