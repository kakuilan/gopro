// 打印金字塔
package main
import "fmt"

func main(){
    //嵌套循环
    for i:=1;i<=100;i++{
	for j:=1;j<=i;j++{
	    fmt.Printf("%s", "A")
	}
	fmt.Println()
    }


    //一个循环
    str := "A"
    for i:=0;i<100;i++{
	fmt.Printf("%s\n", str)
	str = str + "A"
    }
}
