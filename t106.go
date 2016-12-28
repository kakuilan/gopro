//本月月历
package main
import (
    "fmt"
    "time"
)

func main(){
    now := time.Now()
    yr,mo,_ := now.Date()
    d1 := time.Date(yr, mo, 1, 0, 0, 0, 0, time.UTC)
    d2 := time.Date(yr, mo+1, 1, 0, 0, 0, 0, time.UTC)
    d2 = d2.Add(-24 * time.Hour)

    w := d1.Weekday()
    _,ww := d1.ISOWeek()

    fmt.Println("\n周\t日 一 二 三 四 五 六 ")
    fmt.Printf("%d\t%*d", ww, int((w+1)*3), 1)
    for i:=2;i<=d2.Day(); i++{
	if w++; w%7 ==0 {
	    w++
	    fmt.Printf("\n%d\t", ww)
	}
	fmt.Printf("%3d", i)
    }
    fmt.Println("")
}
