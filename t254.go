// 排序接口
package main
import "fmt"

type Sorter interface {
    Len() int
    Less(i,j int) bool
    Swap(i,j int)
}

//定义用于排序slice的新类型
type Xi []int
type Xs []string

//实现Sorter接口的方法
//整数的
func (p Xi) Len() int {return len(p)}
func (p Xi) Less(i,j int) bool {return p[j]<p[i]}
func (p Xi) Swap(i,j int) {p[i],p[j]=p[j],p[i]}
//字符串的
func (p Xs) Len() int {return len(p)}
func (p Xs) Less(i,j int) bool {return p[j]<p[i]}
func (p Xs) Swap(i,j int) {p[i],p[j]=p[j],p[i]}

//用于Sorter接口的通用排序函数
func Sort(x Sorter) {
    for i:=0;i<x.Len()-1;i++{
	for j:=i+1;j<x.Len();j++{
	    if x.Less(i,j) {
		x.Swap(i,j)
	    }
	}
    }

}

func main(){
    ints := Xi{34, 56, 2, 18, 20, 87, 44, 1}
    strings := Xs{"nut", "nil", "age", "element", "zoon", "go"}

    Sort(ints)
    fmt.Println(ints)
    Sort(strings)
    fmt.Println(strings)


}
