//函数作为参数
package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello ", name)
}

func logger(f func(string), name string) {
	fmt.Println("start calling method sayHello:")
	f(name)
	fmt.Println("end calling method sayHello.")
}

func main() {
	logger(sayHello, "中年码农")
}
