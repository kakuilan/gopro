//调用时间函数获取当前日期
package main
import (
		"fmt"
		"time"
		)

func main(){
	//获取时间戳
	t := time.Now()
	
	//按年月日时分秒格式输出
	fmt.Println(t.String())

	//按年月日格式
	fmt.Println(t.Format("206年01月02日"))

	//输出星期几
	fmt.Println(t.Weekday().String())

}
