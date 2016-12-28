// 模仿命令uniq的功能
package main
import "fmt"

func main(){
    list := []string{"a", "b", "a", "c", "a", "d", "e"}
    first := list[0]

    fmt.Printf("%s \n", first)
    for _,v := range list[1:] {
	if first != v {
	    fmt.Printf("%s \n", v)
	    first =v
	}
    }


}
