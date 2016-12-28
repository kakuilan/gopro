package main
import "fmt"

func main(){
    var j int = 5
    //闭包
    a := func()(func()){
	var i int = 10
	return func(){
	    fmt.Printf("i,j: %d, %d\n", i, j)
   	}
    }() //注意要有(),否则无错误无输出

    a()
    j *= 2
    a()

}

