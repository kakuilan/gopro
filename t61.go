package main
import "fmt"

func main(){
    var i uint8
    for i < 8 {
	fmt.Printf("%08b\n", 1<<i)
	i++
    }

    for j:=0;j<8;j++{
	fmt.Printf("%08b\n", 1<<uint(j))
    }
}
