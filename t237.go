// 冒泡排序
package main
import "fmt"

func main(){
    n := []int{5, -1, 0, 12, 3, 6}
    fmt.Println("未排序前:", n)
    bubblesort(n)
    fmt.Println("排序后:", n)
}

func bubblesort(n []int) {
    for i:=0;i<len(n)-1;i++{
	for j:=i+1;j<len(n);j++{
	    if n[j] < n[i] {
		n[i],n[j] = n[j],n[i]
	    }
	}
    }

}
