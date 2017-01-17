//array create
package main
import "fmt"

func main() {
    var buffer [20]byte
    var grid1 [3][3]int
    grid1[1][0], grid1[1][1], grid1[1][2] = 8,6,2
    grid2 := [3][3]int {{4,3}, {8,6,2}}
    cities := [...]string{"Shanghai", "Mumbai", "Istanbul", "Beijing"}
    cities[len(cities)-1] = "Karchi"
    fmt.Println("Type Len Contents")
    fmt.Println(grid1, grid2)
    fmt.Printf("%-8T %2d %v\n", buffer, len(buffer), buffer)
    fmt.Printf("%-8T %2d %v\n", cities, len(cities), cities)

}
