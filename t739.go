//pointer
package main
import "fmt"

func main () {
    i := 9
    j := 5
    product := 0
    swapAndProduct1(&i, &j, &product)
    fmt.Println(i,j, product)

    swapAndProduct2(i,j)
    fmt.Println(i,j,product)
}

func swapAndProduct1(x,y,product *int) {
    if *x > *y {
        *x,*y = *y,*x 
    }
    *product = *x * *y 
}

func swapAndProduct2(x,y int) (int,int,int) {
    if x > y {
        x,y = y,x
    }

    return x,y, x*y
}
