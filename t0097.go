//使用sorter接口排序
package main

import (
	"fmt"
	"gopro/mysort"
)

//整数测试
func ints() {
	data := []int{74, 59, 238, -784, 9845, 9959, 905, 0, 0, 42, 7523, -5467982, 7586}
	a := mysort.IntArray(data) //conversion to type IntArray
	mysort.Sort(a)
	if !mysort.IsSorted(a) {
		panic("int sort fails")
	}
	fmt.Printf("The int sorted array is: %v\n", a)
}

//字符串测试
func strings() {
	data := []string{"monday", "friday", "tuesday", "wednesday", "sunday", "thursday", "", "saturday"}
	a := mysort.StringArray(data)
	mysort.Sort(a)
	if !mysort.IsSorted(a) {
		panic("string sort fail")
	}
	fmt.Printf("The string sorted array is: %v\n", a)
}

//其他类型
type day struct {
	num       int
	shortName string
	longName  string
}
type dayArray struct {
	data []*day
}

func (p *dayArray) Len() int {
	return len(p.data)
}
func (p *dayArray) Less(i, j int) bool {
	return p.data[i].num < p.data[j].num
}
func (p *dayArray) Swap(i, j int) {
	p.data[i], p.data[j] = p.data[j], p.data[i]
}

func days() {
	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturda"}
	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
	a := dayArray{data}
	mysort.Sort(&a)
	if !mysort.IsSorted(&a) {
		panic("day sort fail")
	}
	fmt.Printf("The days sorted array is: %V\n", a)
	for _, d := range data {
		fmt.Printf("%s ", d.longName)
	}
	fmt.Printf("\n")
}

func main() {
	ints()
	strings()
	days()

}
