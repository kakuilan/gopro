// 多个defer注册,按FILO次序执行.哪怕函数或某个延迟调用发生错误,这些调用依旧会被执行.
package main

func test(x int) {
	defer println("a")
	defer println("b")

	defer func(){
		println(100 / x)	//div 0 异常未被捕获,逐步往外传递,最终终止进程
	}()

	defer println("c")
}

func main(){
	test(0)
}
