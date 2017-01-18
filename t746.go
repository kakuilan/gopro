//slice 
package main
import "fmt"

func main () {
    amounts := []float64 {235.22,234.34, 23.66, 366.32, 464.34, 345.67, 234.1, 634.56}
    sum1 := 0.0
    sum2 := 0.0 

    //不改变切片
    for _,amount := range amounts {
        sum1 += amount
    }
    fmt.Printf("E %.1f-> %.1f\n", amounts, sum1)

    //改变切片
    for i := range amounts {
        amounts[i] *= 1.05
        sum2 += amounts[i]
    }

    fmt.Printf("E %.1f-> %.1f\n", amounts, sum2)


}
