package main
import "fmt"

type Err struct{}

func (_ *Err) Error() string {
    return "To err is human"
}

func ToErr(ok bool) error {
    var e *Err = nil
    if ok {
	e = &Err{}
    }
    return e
}

func NoErr (ok bool) error {
    if !ok {
	return &Err{}
    }
    return nil
}

func main(){
    fmt.Println(ToErr(true))
    fmt.Println(ToErr(false))
    fmt.Println(NoErr(true))
    fmt.Println(NoErr(false))
}
