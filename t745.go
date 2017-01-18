//slice 
package main
import "fmt"

func main() {
    bufer := make([]byte, 20, 60)
    grid1 := make([][]int, 3)
    for i:= range grid1 {
        grid1[i] = make ([]int, 3)
    }

    grid1[1][0], grid1[1][1], grid1[1][2] = 8, 6, 2
    grid2 := [][]int{{4,3,0}, {8,6,2}, {0,0,0}}
    cities := []string{"Shanghai", "Mumbai", "Istanbul", "Beijing"}
    cities[len(cities)-1] = "Karachi"
    fmt.Println("Type Len Cap Contents")
    fmt.Printf("%-8T %2d %3d %v\n", bufer, len(bufer), cap(bufer), bufer)
    fmt.Printf("%-8T %2d %3d %v\n", cities, len(cities), cap(cities), cities)
    fmt.Printf("%-8T %2d %3d %v\n", grid1, len(grid1), cap(grid1), grid1)
    fmt.Printf("%-8T %2d %3d %v\n", grid2, len(grid2), cap(grid2), grid2)


}
