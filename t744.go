//切片和其隐藏数组
package main
import "fmt"

func main() {
    
    s := []string{"A","B","C","D","E","F","G"}
    t := s[:5] 
    u := s[3: len(s)-1]
    fmt.Println(s,t,u)
    u[1] = "X"
    fmt.Println(s,t,u)
}
