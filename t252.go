// interface
package main
import "fmt"

type S struct{i int}
func (p *S) Get() int {return p.i}
func (p *S) Put(v int) {p.i=v}

type I interface {
    Get() int
    Put(int)
}

type R struct {i int}
func (p *R) Get() int{return p.i}
func (p *R) Put(v int) {p.i=v}

func f(p I) {
    switch t:=p.(type) {
	case *S:
		fmt.Println(t)
	case *R:
	//case S:
	//case R:
	default:
    }
    fmt.Println(p.Get())
    p.Put(1)
}

func g(something interface{}) int {
    // return something.(I).Get()
    if v,ok := something.(I);ok {//检查是否可以转换
	return v.Get()		//如果可以,调用Get()
    }
    return -1
}



func main(){
    var s S;f(&s)
    var s2 = new(S)
    fmt.Println(g(s2))
    fmt.Println(g(3))
}
