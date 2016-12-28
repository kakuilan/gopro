package main
import (
    "fmt"
    "os"
)

func main(){
    str := "Hello,世界"
    n := len(str)

    for i:=0;i<n;i++{
	fmt.Fprintf(os.Stdout, "%d %v $c\n", i, str[i], str[i]);
    }

    fmt.Fprintf(os.Stdout, "-------------\n")

    for i,v := range str{
	fmt.Fprintf(os.Stdout, "%d %v %c\n", i, v, v)
    }

    fmt.Fprintf(os.Stdout, "--------------\n")

    rs := []rune(str)
    n = len(rs)
    for i := 0;i<n;i++{
	fmt.Fprintf(os.Stdout, "%d %v %c\n", i, rs[i], rs[i])
    }



}
