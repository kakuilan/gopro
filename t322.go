//没有结构化异常,使用panic抛出错误,recover捕获错误
package main

func test(){
	defer func(){
		if err := recover();err!=nil { //这里接收错误
			println("发生错误:", err.(string)) //将 interface{}转型为具体类型
		}
	}()
	panic("panic error!") //这里抛出错误	
}

func main(){
	test()
}
