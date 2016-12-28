// switch fallthrough 继续下一分支
package main

func main(){
	x := 10
	//将会执行2个分支
	switch x{
		case 10:
			println("a")
			fallthrough
		case 0:
			println("b")
	}


}

