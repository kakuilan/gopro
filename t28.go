package main
import "fmt"

func main(){
    sum :=0
    sum2 :=0
    for i:=0;i<10;i++{
	sum += 1
    }

    for{
	sum2++
	if sum2 > 100{
	    break
	}
    }
    fmt.Println(sum, sum2)

}
