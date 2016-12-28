//map应用
//统计今年共几天.若是平年,2月28天,全年共365天;若是闰年,2月29天,全年共366天
//1、定义一个月份天数字典monthdays,string类型的月份名作为键,天数作为值
//2、确定今年是平年还是闰年,若是闰年,将2月份的天数改为29
//3、使用for range 遍历字典monthdays
package main
import (
		"fmt"
		)

func main(){
	monthdays := map[string]int {
			"Jan":31, "Feb":28, "Mar":31,
			"Apr":30, "May":31, "Jun":30,
			"Jul":31, "Aug":31, "Sep":30,
			"Oct":31, "Nov":30, "Dec":31,
		}

	var year string
	fmt.Println("请输入今年是闰年leap还是平年normal...")
	fmt.Scanf("%s", &year)
	if year == "leap" {
		monthdays["Feb"] = 29
	}

	totalDays := 0
	for _,days := range monthdays {
		totalDays += days
	}
	fmt.Printf("今年共%d天\n", totalDays)
}
