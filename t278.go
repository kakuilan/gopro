// 通过自定义类型实现枚举类型限制
package main

type Color int
const (
    Black Color = iota
    Red
    Blue
)

func test(c Color){
    println(c)
}

func main(){
    c := Blue
    test(c)

    x := 1
    //test(x) //发生类型错误
    test(1) //常量会被编译器字段转换
    println("x:",x)


}
