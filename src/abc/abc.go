//在abc包中定义结构体Student

package abc
type Student struct {
    Name string
    Age int
    class string //小写开头的包外不可见,包外的结构体也无法继承该字段
}
