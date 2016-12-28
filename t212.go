// 退出循环
package main
import "fmt"

func main(){
J:  for j:=0;j<5;j++{
	for i:=0;i<10;i++{
	    if i >5 {
		break J //终止J循环,而不是i的那个
		//break //退出i的循环
	    }
	    fmt.Println("j=",j, "i=",i)
	}

	if j>4 {
	    break J
	}
    }

}
