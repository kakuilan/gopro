// interface值
package main
import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

// Human实现Sayhi方法
func (h Human) SayHi(){
	fmt.Println("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//Employee重载Human的SayHi方法
func (e Employee) SayHi(){
	fmt.Println("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

//Interface Men被Human,Student和Employee实现
//因为这3个类型都实现了下面的2个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main(){
	mike := Student{Human{"Mike", 25, "13712345678"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "13998765432"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "13612345679"}, "Golang Inc", 1000}
	tom := Employee{Human{"Tom", 46, "13398765444"}, "Tings Ltd", 5000}

	//定义Men类型的变量i
	var i Men

	//i能存储Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储Employee
	i = tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let`s use a slice of Men and see what happens")
	x := make([]Men, 3)
	//T这3个都是不同类型的元素,但是他们实现了interface融一个接口
	x[0],x[1],x[2] = paul, sam, mike

	for _,value := range x {
		value.SayHi()
	}

}


