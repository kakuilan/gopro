//接口的定义和赋值
package main
import "fmt"

//定义对象People,Teacher,Student
type People struct{
	Name string
}

type Teacher struct {
	People
	Department string
}

type Student struct {
	People
	School string
}

//对象方法实现
func (p People) SayHi(){
	fmt.Printf("Hi,I`m %s.Nice to meet you!\n", p.Name)
}

func (t Teacher) SayHi(){
	fmt.Printf("Hi,my name is %s.I`m working in %s.\n", t.Name, t.Department)
}

func (s Student) SayHi(){
	fmt.Printf("Hi,myname is %s.I`m studing in %s.\n", s.Name, s.School)
}

func (s Student) Study(){
	fmt.Printf("I`m learning Golang in %s.\n", s.School)
}

//定义接口Speaker,Learner
type Speaker interface{
	SayHi()
}

type Learner interface {
	SayHi()
	Study()
}

func main(){
	people := People{"张三"}
	teacher := Teacher{People{"郑五"}, "Computer Science"}
	student := Student{People{"李四"}, "Yale University"}

	var is Speaker //定义Speaker类型的变量is
	is = people //is能存储People
	is.SayHi()

	is = teacher //is能存储Teacher
	is.SayHi()

	is = student //is能存储Student
	is.SayHi()

	var il Learner //定义Learner类型的变量il
	il = student //il能存储Student
	il.Study()

	fmt.Println()

	//因为接口is能够同时持有3种类型的对象,所有可定义一个包含Speaker接口类型元素的Slice
	ix := make([]Speaker, 3)
	ix[0],ix[1],ix[2] = people,teacher,student
	for _,value := range ix {
		value.SayHi()
	}

}
