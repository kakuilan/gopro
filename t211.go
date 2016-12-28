// goto 该标签大小写敏感
package main
import "fmt"

func myfunc(){
    i := 0
Here:
    fmt.Println(i)
    if(i >100){
	return
    }
    i++
    goto Here
}

func main(){
    myfunc()
}
