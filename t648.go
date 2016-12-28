//Method的重写
package main
import "fmt"

type people struct {
	name string
	phone string
}

type teacher struct {
	people
	department string
}

type student struct {
	people
	school string
}

func (r people) sayHi(){
	fmt.Printf("Hi,I`m %s .you can call me on %s.\n", r.name, r.phone)
}

//重写方法
func (s student) sayHi(){
	fmt.Printf("Hi,I`m %s,I study in %s,call me on %s.\n", s.name, s.school, s.phone)
}

func main(){
	t1 := teacher{people{"张三","010-222001"},"Computer Science"}
	s1 := student{people{"李四","020-110032"},"Yale University"}
	t1.sayHi()
	s1.sayHi()


}
