//子类继承
package main

import "fmt"

type O struct {
    h int
}

func (a *O) Get() int {
    return a.h
}

func (a *O) Set(h int) {
    a.h = h
}

type OO struct {
    O
    i int
}

func (a *OO) Get() int {
    a.i++
    return a.O.Get()
}

func main(){
    oo := new(OO)
    oo.Set(42)
    fmt.Println(oo, oo.Get())
}
