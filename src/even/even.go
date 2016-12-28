// even包演示
package even

//可导出的函数
func Even(i int) bool {
    return i % 2 ==0
}

//私有函数
func odd(i int) bool {
    return i % 2 ==1
}
