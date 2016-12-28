// 常量
package main

const x,y int = 1,2
const s = "Hello, World!"
const (
    a,b = 10,100
    c bool = false
    d = len(s)
)

func main(){
    const x = "xxx" //未使用的局部常量不会引发编译错误
    println(s,a,b,c)
}
