//映射
package main
import "fmt"

func main() {
    massForPlanet := make(map[string]float64) //与map[string]float64{}相同
    massForPlanet["Mercury"] = 0.06
    massForPlanet["Venus"] = 0.82
    massForPlanet["Earth"] = 1.00
    massForPlanet["Mars"] = 0.11
    fmt.Println(massForPlanet)


}
