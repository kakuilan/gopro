package main
import "fmt"

func Add(x,y int){
    z := x+y
    fmt.Println(x,y,z)
}

func main(){
    for i:=0;i<10;i++{
	go Add(i, i)
    }
    var input string
    fmt.Scanln(&input)
}
