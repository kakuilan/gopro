//结构体的方法
package main
import "fmt"

type Student struct {
    Name string
    Age int
    class string
}

func (this Student) getName() string {
    return this.Name
}

//结构体可以传指针类型
func (this *Student) getAge() int {
    return this.Age
}

func main(){
    s := Student{Name: "张三", Age : 15, class : "44班"}
    fmt.Println(s.getName(), s.getAge())

}
