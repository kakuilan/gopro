//通过channel实现同步机制
//在main函数中起一个goroutine,通过非缓冲队列的使用,能够保证在goroutine执行结束之前main函数不会提前退出

package main

func worker(done chan bool) {
	println("start working...")
	done <- true
	println("end working...")
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	//注释掉下句,将不会等待worker执行完就退出
	<-done
}
