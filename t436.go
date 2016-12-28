// goto.用goto跳转必须在当前函数内定义的标签.标签名是大小写敏感的.
package main

func myFunc() int {
	i := 0
Here: //这行的第一个词,以冒号结束作为标签
	println(i)
	if i >10 {
		return i
	}
	i++
	goto Here //跳转到Here去
}

func main(){
	myFunc()
}
