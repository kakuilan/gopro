//struct结构体
package main
import "fmt"

type Student struct {
    Name string
    Age int
    class string
}

func main(){
    s1 := new(Student)
    s1.Name = "张三"
    s1.Age = 12
    s1.class = "22班"
    s2 := Student{"李四",22,"21班"}
    fmt.Println(s1,s2)
}

