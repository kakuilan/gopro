package main
import "fmt"

func main(){
    s := [4]int{0,1,2,3}
    t := s[1:3]
    fmt.Println(s, t)
    fmt.Println(len(s), cap(s))
    fmt.Println(len(t), cap(t))
}
