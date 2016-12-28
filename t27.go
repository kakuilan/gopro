package main
import "fmt"

func main(){
    Num := 5
    switch {
	case 0<= Num && Num <=3:
	    fmt.Println("0~3")
	case 4<=Num && Num<=6:
	    fmt.Println("4~6")
	case 7<=Num && Num<=9:
	    fmt.Println("7~9")
    }

}
