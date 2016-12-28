// 求平均数
package main
import "fmt"

func average(xs []float64) (avg float64) {
    sum := 0.0
    switch len(xs) {
	case 0:
	    avg = 0
	default:
	    for _,v := range xs {
		sum += v
	    }
	    avg = sum / float64(len(xs))
    }

    return
}

func main(){
    nums := []float64{1.0, 2.22, 4.5, 55.6}
    avg := average(nums)
    fmt.Println(avg)
}
