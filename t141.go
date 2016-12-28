//结构体继承

package main
import "fmt"

type Student struct {
    Name string
    Age int
    class string
}

//结构体可以传指针类型
func (this *Student) Display() {
    fmt.Println(this.Name, ",", this.Age)
}

//定义一个大学生类,继承Student
//CollegeStudent将继承Student的所有字段和方法
type CollegeStudent struct {
    Student
    Profession string
}

//CollegeStudent也可以重写继承的方法
func (this *CollegeStudent) Display(){
    fmt.Println("重写的方法:", this.Name,",",this.Age,",",this.Profession)
}

func main(){
    s1 := CollegeStudent {
	Student : Student{Name:"李四", Age:25, class :"2004(3)班"},
	Profession :"物理",
    }
    s1.Display()
    fmt.Println(s1.Student.Name) //可以通过student访问Name
    fmt.Println(s1.Name) //也可以直接通过Name访问

}
