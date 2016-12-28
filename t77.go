package main
import (
    "fmt"
    "math"
    "sort"
)

func main(){
    is := sort.IntSlice{3, 1, 4,1, 5, 8,2,6}
    ss := sort.StringSlice{"是","啊","A","3","b","_","商"}
    fs := sort.Float64Slice{math.Inf(-1), math.Inf(+1), math.NaN()}

    sort.Sort(is)
    sort.Sort(ss)
    sort.Sort(fs)
    fmt.Println(is, ss, fs)

}
