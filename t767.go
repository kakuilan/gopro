//map test
package main
import (
    "fmt"
)

func main() {
    populationForCity := map[string]int{
        "Istanbul":1261000,
        "Karochi":10620000, "Mumbai":12690000, "Shanghai":13680000,
    }


    cityForPopulation := make(map[int]string, len(populationForCity))
    for city,population := range populationForCity {
        cityForPopulation[population] = city
    }
    fmt.Println(cityForPopulation)

}
