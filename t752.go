//切片排序
package main
import (
    "fmt"
    "strings"
    "sort"
)

func SortFoldedStrings(slice []string) {
    sort.Sort(FoldedStrings(slice))
}

type FoldedStrings []string
func (slice FoldedStrings) Len() int {
    return len(slice)
}
func (slice FoldedStrings) Less(i,j int) bool {
    return strings.ToLower(slice[i]) < strings.ToLower(slice[j])
}
func (slice FoldedStrings) Swap(i,j int) {
    slice[i],slice[j] = slice[j],slice[i]
}



func main() {
    files := []string{"Test.conf","uitl.go","Makefile","misc.go","main.go"}
    fmt.Printf("Unsorted: %q\n", files)
    sort.Strings(files) //标准库中的排序函数
    fmt.Printf("Underlying bytes: %q\n", files)
    SortFoldedStrings(files) //自定义排序函数
    fmt.Printf("Case insensitive: %q\n", files)

}
