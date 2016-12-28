package main
import "fmt"

type Point struct {
    x, y float64
}

type Rect struct {
    min, max Point
}

func (r *Rect) Area() float64 {
    w := r.max.x - r.min.x
    h := r.max.y - r.min.y
    return w * h
}

func (r Point) Area() float64 {
    return 0
}

func main(){
    var r Rect
    r.max = Point{2, 2}
    fmt.Println(r.Area(), (&r).Area(), (*Rect).Area(&r))
    p := &r.min
    fmt.Println(p.Area(), (*p).Area(), Point.Area(r.min))
}
