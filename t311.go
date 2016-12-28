// 多返回值可直接作为其他函数调用实参
package main

func test()(int,int) {
	return 1,2
}

func add(x,y int) int {
	return x+y
}

func sum(n ...int) int {
	var x int
	for _,i := range n{
		x +=i
	}
	return x
}

func main(){
	println(add(test()))
	println(sum(test()))

}


