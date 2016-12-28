//打印中文日期信息,并使用switch判断当前是星期几,打印出来
//1.使用time.Now()函数获取系统当前时间戳
//2.使用时间对象的Format()方法,按照"年月日"格式输出中文日期信息
//3.使用时间对象的Weekday()方法获取星期信息,再使用String()方法将它转换为字符串格式
//4.使用switch判断星期
package main
import (
		"fmt"
		"time"
		)

func main(){
	//获取时间戳
	t := time.Now()
	//按年月日格式输出日期信息
	fmt.Printf(t.Format("2006年01月02日"))
	switch t.Weekday().String() {
		case "Sunday":
			fmt.Println(" 星期天")
		case "Monday":
			fmt.Println(" 星期一")
		case "Tuesday":
			fmt.Println(" 星期二")
		case "Wednesay":
			fmt.Println(" 星期三")
		case "Thursday":
			fmt.Println(" 星期四")
		case "Friday":
			fmt.Println(" 星期五")
		case "Saturday":
			fmt.Println(" 星期六")
	}

}
