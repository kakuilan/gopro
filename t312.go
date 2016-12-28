// 命名返回参数可看做与形参类似的局部变量,最后由return 隐式返回
package main

func add(x,y int) (z int) {
	z = x+y
	return
}

func main(){
	println(add(1,2))

}
