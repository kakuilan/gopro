package main
import "fmt"

type Stringer interface{
    String() string
}

func Tprintln(args ...interface{}){
    for _, arg := range args{
	switch t := arg.(type){
	    case int:
	    case string:
	    default:
		if v, ok:= arg.(Stringer); ok{
		    val := v.String()
		    fmt.Println(t, "String:", val)
		}else{
		    fmt.Println("default:", t)
		}
	}
    }
}

func main(){
    var ns Stringer;
    Tprintln(12,"asdf", 12323.33, ns) 
}
