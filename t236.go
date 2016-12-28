// slice最大最小值
package main
import "fmt"

func max(l []int) (max int) {
    max = l[0]
    for _,v := range l {
	if v > max {
	    max = v
	}
    }
    return
}

func min(l []int) (min int) {
    min = l[0]
    for _,v := range l{
	if v < min {
	    min = v
	}
    }

    return
}

func main(){
    sl := []int{1,3,45,4,676,2,8}
    max := max(sl)
    min := min(sl)

    fmt.Println(max, min)
}
