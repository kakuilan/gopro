// struct
package main
import "fmt"

type person struct {
	name string
	age int
}



func main(){
	var p1 person //p1现在是person类型的变量了
	p1.name = "kkl" //赋值给p的name属性
	p1.age = 23
	fmt.Printf("The p1 person`s name is %s \n", p1.name) //访问p1的name属性

	//按照顺序提供初始化值
	p2 := person{"Tom", 24}

	//通过field:value的方式初始化
	p3 := person{age:34, name:"yyt"} 
	fmt.Println(p2, p3)

}
