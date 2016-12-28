// 在函数内goto跳转,签名区分大小写
package main

func main(){
	var i int
	for {
		println(i)
		i++
		if i>2 {goto BREAK}
	}

BREAK:
	println("break")

//EXIT: // error: label EXIT defined and not used

}


