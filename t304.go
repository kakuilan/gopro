// switch 省略条件表达式
package main

func main(){
	x := []int {1,2,3,4}
	
	switch {
		case x[1]>0:
			println("a")
		case x[1]<0:
			println("b")
		default:
			println("c")
	}

	//带初始化语句
	switch i := x[2]; {
		case i>0:
			println("a")
		case i<0:
			println("b")
		default:
			println("c")
	}

}
