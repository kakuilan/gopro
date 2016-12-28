// 复合声明
package main
import "fmt"

func main(){
    ar := [...]string{Enone:"no error", Einval:"invalid argument"}
    sl := []string{Enone: "no error", Einval:"invalid argument"}
    mp := map[int]string {Enone:"no error", Einval:"invalid argument"}
    fmt.Println(ar, sl, mp)

}
