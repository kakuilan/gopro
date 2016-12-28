// 匿名函数可赋值给变量,作为结构字段,或者在channel里传送
package main

func main(){
	fn := func(){println("Hello, World!")}
	fn()

	//函数列表
	fns := [](func(x int) int) {
		func(x int) int {return x+1},
		func(x int) int {return x+2},
	}

	println(fns[0](100))

	//函数作为字段
	d := struct {
		fn func() string
	}{
		fn: func() string {return "Hello, World! Second!"},
	}

	println(d.fn())

	//传送函数的channel
	fc := make(chan func()string, 2)
	fc <- func()string{return "Hello, Channel function!"}
	println((<-fc)())
}
