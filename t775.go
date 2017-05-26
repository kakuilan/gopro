//将函数调用作为函数的参数
package main
import (
    "fmt"
    "math"
)

func Heron(a,b,c int) float64 {
    a1,b1,c1 := float64(a),float64(b),float64(c)
    s := (a1+b1+c1)/2
    return math.Sqrt(s *(s-a1)*(s-b1)*(s-c1)) 
}

func PythagoreanTriple(m,n int) (a,b,c int) {
    if m < n {
        m,n = n,m
    }
    return (m*m)-(n*n),(2*m*n),(m*m)+(n*n)
}

func main() {
    for i:=1;i<=4;i++{
        a,b,c := PythagoreanTriple(i,i+1)
        s1 := Heron(a,b,c)
        s2 := Heron(PythagoreanTriple(i,i+1))
        fmt.Printf("s1 == %10f == s2 == %10f\n", s1, s2)
    }


}
