package main
import "fmt"

type Duck float32
func (Duck) Quack() string {
    return "ga"
}

type Goose struct {Duck}
type Quacker interface {
    Quack() string
}

func main(){
    var d Duck = 0
    var q Quacker
    g := Goose{d}
    g.Quack()
    q = g
    fmt.Println(q.Quack())
}
