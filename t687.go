//goroutine
package main
import "fmt"

func test(i int) {
	fmt.Println("Go...")
	fmt.Println(i)
}

func main(){
	for i:= 0;i<10;i++{
		go test(i)
	}

	fmt.Println("main run...")
}

