package main
import (
    "fmt"
    "sort"
)

type Rev struct {
    sort.Interface
}

func (r Rev) Less(i, j int) bool {
    return r.Interface.Less(j, i)
}

func main(){
    a := []int{6, 7, 4, 2}
    sort.Ints(a)
    fmt.Println(a)
    sort.Sort(Rev{sort.IntSlice(a)})
    fmt.Println(a)
}
