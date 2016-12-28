//接口的转换
package main
import "fmt"

type People struct {
	Name string
	Age int
}

type Student struct {
	People
	School string
}

type PeopleInfo interface {
	GetPeopleInfo()
}

type StudentInfo interface {
	GetPeopleInfo()
	GetStudentInfo()
}

func (p People) GetPeopleInfo(){
	fmt.Println("peopleinfo:", p)
}

func (s Student) GetStudentInfo(){
	fmt.Println("sutdentinfo:", s)
}

func main(){
	var is StudentInfo = Student{People{"李四",18}, "Yale University"}
	is.GetStudentInfo()
	is.GetPeopleInfo()

	var ip PeopleInfo = is //超集转换为子集
	ip.GetPeopleInfo()
	fmt.Println(is,ip)
}


