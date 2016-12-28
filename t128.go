//for

package main
import "fmt"

func main(){
    str := "Hello world."
    for i,j :=0,len(str);i<j;i++ {
	fmt.Println(string(str[i]))
    }

    i :=0
    for {
	if i >10 {
	    break
	}
	fmt.Println(i)
	i++
    }
}
