//slice
package main
import "fmt"

func main(){
	var array = [10]byte {'a','b','c','d','e','f','g','h','i','j'}
	var aSlice,bSlice []byte

	aSlice = array[:3] //=array[0:3]
	fmt.Println(aSlice)

	aSlice = array[5:] //=array[5:10]
	fmt.Println(aSlice)

	aSlice = array[:] //=array[0:10]
	fmt.Println(aSlice)

	aSlice = array[3:7]
	bSlice = aSlice[3:7]
	fmt.Println(bSlice, len(bSlice), cap(bSlice))
	bSlice = aSlice[1:3]
	fmt.Println(bSlice, len(bSlice), cap(bSlice))
	bSlice = aSlice[:3]
	fmt.Println(bSlice, len(bSlice), cap(bSlice))
	bSlice = aSlice[0:5]
	fmt.Println(bSlice, len(bSlice), cap(bSlice))
	bSlice = aSlice[:]

	fmt.Println(bSlice, len(bSlice), cap(bSlice))

}



