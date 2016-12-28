// map
package main
import "fmt"

func main(){
    monthdays := map[string]int{
	"Jan":31, "Feb":28, "Mar":31,
	"Apr":30, "May":31, "Jun":30,
	"Jul":31, "Aug":31, "Sep":30,
	"Oct":31, "Nov":30, "Dec":31, //逗号是必须的
    }

    //使用make形式
    map1 := make(map[string]int)
    fmt.Println(map1)

    yeardays := 0
    for _,days := range monthdays {
	yeardays += days
    }
    fmt.Printf("Numbers of days in a year: %d\n", yeardays)

    //向map添加元素
    monthdays["Undecim"] = 30//添加一个月
    monthdays["Feb"] = 29 //闰年时重写这个元素


    //检查元素是否存在
    var value int
    var present bool
    value,present = monthdays["Jan"] //如果存在,present值为true
    v,ok := monthdays["Jan"]
    fmt.Println(value,present,v,ok)

    //从map移除元素
    delete(monthdays, "Mar") //删除"Mar"元素
    fmt.Println(monthdays)


}
