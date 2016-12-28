// 函数命名返回参数允许defer延迟调用通过闭包读取和修改
package main

func add(x,y int) (z int) {
	defer func(){
		z += 100
	}()

	z = x+y
	return
}

func main(){
	println(add(1,2)) //输出103
}

