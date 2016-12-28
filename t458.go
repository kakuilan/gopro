// channel
package main
import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _,v := range a{
		sum += v
	}
	c <- sum //send sum to c
}

func main(){
	a := []int{25,23, 6, -9,3,0}
	
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x,y := <-c, <-c //recive from c
	fmt.Println(x,y, x+y)

}
