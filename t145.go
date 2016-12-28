//接口
package main
import "fmt"

type Student struct {
    Name string
    Age int
    class string
}

type IStudent interface {
    GetName() string
    GetAge() int
}

func (this *Student) GetName() string {
    return this.Name
}

func (this *Student) GetAge() int {
    return this.Age
}

func main(){
    var s1 IStudent = &Student{"李四",23,"2003(4)班"}
    fmt.Println(s1.GetName())


}
