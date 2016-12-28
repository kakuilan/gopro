//在main包中继承abc包的Student

package main
import (
    "abc"
    "fmt"
)

type MyStudent struct {
    abc.Student
}

func main(){
    s := MyStudent{}
    s.Student.Name = "karl"
    //下句将报错
    //s.Student.class = "4班"
    fmt.Println(s)
}
