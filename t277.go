// iota打断与恢复
package main

const (
    A = iota
    B
    C = "C"
    D
    E = iota
    F
)

func main(){
    println(A,B,C,D,E,F)

}
