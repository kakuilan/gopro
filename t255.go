// 通用的计算最大值
package main
import "fmt"

func Less(l, r interface{}) bool {
    switch l.(type) {
	case int :
	    if _,ok := r.(int);ok {
		return l.(int) < r.(int)
	    }
	case float32 :
	    if _,ok := r.(float32);ok {
		return l.(float32) < r.(float32)
	    }
    }
    return false
}

func main(){
    var a,b,c int = 5,24,0
    var x,y,z float32 = 2.4, 43.21, 0.0

    if c = a;Less(a,b){
	c = b
    }
    if z=x;Less(x,y) {
	z = y
    }

    fmt.Println(c,z)

}
