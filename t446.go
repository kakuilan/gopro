// 函数作为类型
package main
import "fmt"

type testInt func(int) bool //声明一个函数类型

func isOdd(integer int) bool {
	if integer %2 ==0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer %2 == 0 {
		return true
	}
	return false
}

//声明的函数类型在这里当做一个参数
func filter(slice []int, f testInt) []int {
	var result []int
	for _,value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main(){
	slice := []int {1,2,3,4,5,6,7}
	fmt.Println("slice = ", slice)

	odd := filter(slice, isOdd) //函数当做值来传递
	fmt.Println("是偶数的元素: ", odd)

	even := filter(slice, isEven)
	fmt.Println("是奇数的元素: ", even)


}

