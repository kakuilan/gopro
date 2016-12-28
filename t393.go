// 初始化函数.
// 因为无法保证初始化函数执行顺序,因此全局变量应该直接用var初始化
package main
import (
		"time"
		"fmt"
		)

var now = time.Now()

func init(){
	fmt.Printf("Now: %v\n", now)
}

func init(){
	fmt.Printf("since:%v\n", time.Now().Sub(now))
}

func main(){
	fmt.Println("初始化函数测试，哈哈哈！")
}
