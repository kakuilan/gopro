//使用new获取指针
package main
import "fmt"

type composer struct {
    name string
    birthYear int
}

func main() {
    antonio := composer{"antonio Teixiera", 1707} //composer类型值
    agnes := new(composer) //指向composer的指针
    agnes.name, agnes.birthYear = "Agnes Zimmermann",1845 
    julia := &composer{} //指向composer的指针 
    julia.name, julia.birthYear = "Julia Word Howe",1819
    augusta := &composer{"Augusta Holmes", 1847} //指向composer的指针
    fmt.Println(antonio)
    fmt.Println(agnes, augusta, julia)

}
