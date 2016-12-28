//显式return 返回前,会先修改命名返回参数
package main

func add(x,y int) (z int) {
	defer func(){
		println(z) //输出 203
	}()

	z = x+y	//执行顺序: (z=z+200) -> (calll defer) -> (ret)
	println(z)
	return z+200
}

func main(){
	println(add(1,2)) //输出 203
}
