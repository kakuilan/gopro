//键为指针类型的map
package main
import "fmt"

type Point struct {x,y,z int}
func (point Point) String() string {
    return fmt.Sprintf("(%d,%d,%d)", point.x, point.y, point.z)
}

func main() {
    //指针作为键
    triangle := make(map[*Point]string, 3)
    triangle[&Point{89,47,27}] = "a"
    triangle[&Point{85,65,44}] = "b"
    triangle[&Point{7,22,35}] = "c"
    fmt.Println(triangle)

    //结构体作为键
    nameForPoint := make(map[Point]string) //等同于map[Point]string{}
    nameForPoint[Point{54,91,78}] = "x"
    nameForPoint[Point{54,12,24}] = "y"
    fmt.Println(nameForPoint)

    //遍历
    populationForCity := map[string]int{"Istanbul":1261000,
    "Karochi":10620000, "Mumbai":12690000, "Shanghai":13680000}

    for city,population := range populationForCity {
        fmt.Printf("%-10s %8d\n",city,population)
    }

}
