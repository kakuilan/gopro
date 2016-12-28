//switch
package main
import "fmt"

func main(){
    i := 5
    switch i {
	case 1:
	    fmt.Println("i is equal to 1")
	case 2:
	    fmt.Println("i is equal to 2")
	case 3,4,5,6: //case可以有多个值
	    fmt.Println("i is equal to 3,4,5 or 6")
	    //fallthrough //添加后,相当于去掉break,继续执行下面
	default:
	    fmt.Println("others")
    }

    result := 0
    switch {
	case result<0 :
	    fmt.Println("result小于0")
	case result >0 :
	    fmt.Println("result大于0")
	default :
	    fmt.Println("result等于0")
    }


}
