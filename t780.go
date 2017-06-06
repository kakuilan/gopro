//自定义类型:重写方法
package main
import (
    "fmt"
)

type Item struct {
    id  string //具名字段,聚合
    price float64 //具名字段,聚合
    quantity int //具名字段,聚合
}

func (item *Item) Cost() float64 {
    return item.price * float64(item.quantity)
}

type SpecialItem struct {
    Item //匿名字段,嵌入
    catalogId int //具名字段,聚合
}

type LuxuryItem struct {
    Item //匿名字段,嵌入
    markup float64 //具名字段,聚合
}
//重写Item.Cost方法
func (item *LuxuryItem) Cost() float64 {
    return item.Item.Cost() * item.markup
}



func main() {
    special := SpecialItem{Item{"Green", 3, 5}, 207}
    fmt.Println(special.id, special.price, special.quantity, special.catalogId)
    fmt.Println(special.Cost())

}

