//channel遍历和关闭

package main

func main() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	res := 0
	for v := range ch {
		println(v)
		res += v
	}

	println(res)

}
