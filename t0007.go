//自定义类型
package main

import "fmt"

type City string
type Age int

func printAge0(age int) {
	fmt.Println("Age0 is ", age)
}

func printAge1(age Age) {
	fmt.Println("Age1 is ", age)
}

func main() {
	city := City("北京")
	fmt.Println("I live in ", city+" 上海") //字符串拼接
	fmt.Println(len(city))                //len方法

	//自定义类型的原始类型的所有操作同样适用
	middle := Age(12)
	if middle >= 12 {
		fmt.Println("Middle is bigger than 12.")
	}
	printAge0(int(middle))
	printAge1(middle)

}
