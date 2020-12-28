//数组的声明
package main

import "fmt"

func main() {
	//正常声明
	var arrAge = [5]int{18, 20, 15, 22, 16}

	//只声明前面的N个元素
	var arr2 = [10]int{1, 2, 3}

	//未知数组长度
	var arrLazy = [...]int{5, 6, 7, 8, 9, 22}

	//键值对语法
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}

	fmt.Println(arrAge)
	fmt.Println(arr2)
	fmt.Println(arrLazy)
	fmt.Println(arrKeyValue)

}
