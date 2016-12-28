// switch 字节数组比较
package main
import "fmt"

func Compare(a,b []byte) int {
    for i:=0;i<len(a) && i<len(b);i++{
	switch{
	    case a[i] > b[i] :
		return 1
	    case a[i] < b[i] :
		return -1
	}
    }

    switch {
	case len(a) < len(b) :
	    return -1
	case len(a) > len(b) :
	    return 1
    }
    return 0
}

func main(){
    s := "abc"
    a := []byte(s)
    b := []byte("abd")
    c := Compare(a,b)
    fmt.Println(a, b, c)
}
