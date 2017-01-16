//指针
package main
import (
    "fmt"
    "image/color"
)

type rectangle struct {
    x0,y0, x1,y1 int
    fill color.RGBA
}

func main () {
    rect := rectangle{4,8,20,30, color.RGBA{1,2,3,4}}
    fmt.Println(rect)
    resizeRect(&rect, 5,5)
    fmt.Println(rect)

}

func resizeRect(rect *rectangle, width, height int) {
    (*rect).x1 += width //显式地解引用
    rect.y1 += height //.操作符能够自动解引用结构体
}
