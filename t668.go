//接口
package main
import "fmt"

type People struct {
	Name string
}

type Student struct {
	People
	School string
}

type Teacher struct {
	People
	Department string
}

func (p People) SayHi(){}
func (s Student) SayHi(){}
func (t Student) SayHi(){}
func (s Student) Study(){}

type Speaker interface{
	SayHi()
}

type Learner interface{
	SayHi()
	Study()
}

//接口组合
type SpeakLearner interface{
	Speaker
	Learner
}

func main(){

}
