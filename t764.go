//murmur3测试
package main
import (
    "murmur3"
    "fmt"
)

func main() {
    key := []byte("this is a rand test key")
    num1 := murmur3.Sum32(key)
    num2 := murmur3.Sum64(key)
    num3,num4 := murmur3.Sum128(key)
    fmt.Println(num1, num2, num3, num4)

}
