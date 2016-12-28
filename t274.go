//重新赋值与定义新同名变量的区别
package main

func main(){
    s := "abc"
    println(&s)

    s,y := "hello",20 //s被重新赋值
    println(&s, y)

    {
	s,z := 1000, 30 //不在同一层次代码块,定义新同名变量
	println(&s, z)
    }


}
