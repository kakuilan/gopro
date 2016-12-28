// 命名返回参数可被同名局部变量遮蔽,此时需要显示返回
package main

func add(x,y int) (z int) {
	{
		var z = x+y
		//return //Error: z is shadowed during return
		return z //必须显式返回
	}
}

func main(){
	a := add(2,34)
	println(a)
}


