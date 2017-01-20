//切片元素删除
package main
import "fmt"

func main() {
    s1 := []string{"A", "B", "C", "D", "E", "F", "G"}
    s1 = s1[2:] //从开始处删除s[:2]子切片
    fmt.Println(s1)

    s2 := []string{"A", "B", "C", "D", "E", "F", "G"}
    s2 = s2[:4] //从末尾处删除s[4:]子切片
    fmt.Println(s2)

    s3 := []string{"A", "B", "C", "D", "E", "F", "G"}
    s3 = append(s3[:1], s3[5:]...) //从中间删除s[1:5]
    fmt.Println(s3)

}
