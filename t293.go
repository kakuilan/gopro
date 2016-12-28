// 用type在全局或函数内定义新类型
package main

func main(){
    type bigint int64
    var x bigint = 100
    println(x)

}
