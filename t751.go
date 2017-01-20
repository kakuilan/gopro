//切片元素删除
package main
import "fmt"

func RemoveStringSliceCopy(slice []string, start, end int) []string {
    result := make([]string, len(slice)-(end-start))
    at := copy(result, slice[:start])
    copy(result[at:], slice[end:])
    
    return result
}

func main() {
    s := []string{"A","B","C","D","E","F","G"}
    x := RemoveStringSliceCopy(s,0,2) //从头部删除s[:2]
    y := RemoveStringSliceCopy(s,1,5) //从中间删除s[1:5]
    z := RemoveStringSliceCopy(s,4,len(s)) //从结尾删除s[4:]
    fmt.Printf("%v\n%v\n%v\n%v\n", s,x,y,z)


}
