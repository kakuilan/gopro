//for 
package main
import "fmt"

func main(){
oute:
    for i:=0;i<5;i++{
	for k:=0;k<100;k++{
	    if k>1 {
		continue oute
	    }
	    fmt.Println("i=",i, "k=",k)
	}
    }

}
