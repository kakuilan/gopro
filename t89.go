package main

import (
    "fmt"
    "time"
)


func sheep(i int){
    for ;;i +=2 {
	fmt.Println(i, "只羊")
    }
}

func main(){
    go sheep(1)
    go sheep(2)
    time.Sleep(time.Millisecond)
}
