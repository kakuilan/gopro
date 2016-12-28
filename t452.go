// method继承
package main
import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human //匿名字段
	school string
}

type Employee struct {
	Human
	company string
}

//在Human上面定义一个method
func (h *Human) SayHi(){
	fmt.Printf("Hi, I am %s you can cal me on %s\n", h.name, h.phone)
}

//Employee的method重写Human的method
func (e *Employee) SayHi(){
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
			e.company, e.phone)
}

func main(){
	mark := Student{Human{"Mark", 25, "13912345678"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "13898765432"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()

}
