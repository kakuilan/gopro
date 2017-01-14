//修改切片
package main
import "fmt"

func main () {
    grades := []int {87, 55, 43, 71, 60, 43, 32, 19, 64}
    fmt.Println(grades) 
    inflate(grades, 3)
    fmt.Println(grades)

}

func inflate(numbers []int, factor int) {
    for i:= range numbers {
        numbers[i] *= factor
    }
}
