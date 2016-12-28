// 延迟调用参数在注册时求值或复制,可用指针或闭包"延迟"读取
package main

func test(){
	x,y := 10, 20
	
	defer func(i int){
		println("defer:", i, y) // y闭包引用
	}(x)	//x被复制

	x += 10
	y += 100
	println("x =", x, " y=",y)
}

func main(){
	test()
}
