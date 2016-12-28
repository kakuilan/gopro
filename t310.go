// 不能用容器对象接收函数的多返回值.只能用多个变量,或"_"忽略
func test() (int,int) {
	return 1,2
}

func main(){
	//下面是错误的
	// s:=make([]int, 2)
	// s = test()

	//正确的
	x,_ := test()
	println(x)


}



