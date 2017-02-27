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

    population2 := populationForCity["Mumbai"]
    fmt.Println("Mumbai`s population is", population2)
    population2 = populationForCity["Emerald City"]
    fmt.Println("Emerald City`s population is", population2)

    //判断键是否存在
    city := "Istanbul"
    if population,found := populationForCity[city]; found {
        fmt.Printf("%s`s population is %d\n", city, population)
    }else{
        fmt.Printf("%s`s population data is unavailable\n", city)
    }

    city = "Emeriald City"
    _,present := populationForCity[city]
    fmt.Printf("%q is in the map == %t\n", city, present)

    //删除
    fmt.Println(len(populationForCity), populationForCity)
    delete(populationForCity, "Shanghai")
    fmt.Println(len(populationForCity), populationForCity)

    //更新
    populationForCity["Karachi"] = 11620000
    fmt.Println(len(populationForCity), populationForCity)

    //插入
    populationForCity["Beijing"] = 1129000
    fmt.Println(len(populationForCity), populationForCity)


}
