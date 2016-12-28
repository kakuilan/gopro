//从string到byte/int slice的转换
package main
import "fmt"

func main(){
	var str string = "Golang"
	var byteSlice []byte
	byteSlice = []byte(str)

	for _,v := range byteSlice {
		fmt.Printf(" %c \n", v)
	}

}
