//嵌套循环
package main
import "fmt"

func main(){
oute:
    for i:=0;i<5;i++{
	for k:=0;k<5;k++ {
	    if i>0 {
		break oute
	    }
	    fmt.Println("k=",k)
	}	



    }



}
