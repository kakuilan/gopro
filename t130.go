//for
package main
import "fmt"

func main(){
    str := "Hellow world"
    for i,j :=0,len(str);i<j;i++{
	if string(str[i]) ==" " {
	    continue
	}
	fmt.Println(string(str[i]))
    }

}
