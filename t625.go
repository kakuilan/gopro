//panic-recover机制
package main
import "fmt"

func main(){
	var i int
	fmt.Println("请输入一个整数...")
	fmt.Scan(&i)
	
	//必须要先声明defer,否则捕获不到panic异常
	defer func(){
		fmt.Println("函数defer开始运行...")
		if err := recover();err !=nil {
			//这里的err就是panic传入的内容
			fmt.Println("程序异常退出:",i, err)
		}else{
			fmt.Println("程序正常退出.", i)
		}
	}()

	f(i)
}

func f(a int){
	fmt.Println("函数f开始运行...")
	if a > 100{
		panic("参数值超出范围!")
	}else{
		fmt.Println("函数f调用结束.")
	}
}
