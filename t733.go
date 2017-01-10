//strconv
package main
import (
    "fmt"
    "strconv"
)

func main() {
    x,err := strconv.ParseFloat("-99.7", 64)
    fmt.Printf("%8T %6v %v\n", x,x,err)
    y,err := strconv.ParseInt("71399", 10, 0)
    fmt.Printf("%8T %6v %v\n", y,y,err)
    z,err := strconv.Atoi("71399")
    fmt.Printf("%8T %6v %v\n", z,z,err)

    s := strconv.FormatBool(z >100)
    fmt.Println(s)

    i,err := strconv.ParseInt("0xDEED", 0, 32)
    fmt.Println(i,err)

    j,err := strconv.ParseInt("0707", 0, 32)
    fmt.Println(j,err)

    k,err := strconv.ParseInt("10111010001", 2, 32)
    fmt.Println(k,err)

    m := 16769023
    fmt.Println(strconv.Itoa(m))
    fmt.Println(strconv.FormatInt(int64(m), 10))
    fmt.Println(strconv.FormatInt(int64(m), 2))
    fmt.Println(strconv.FormatInt(int64(m), 16))

    t := "Alle ~!``nsker #！#a vare fri."
    quoted := strconv.Quote(t)
    fmt.Println(quoted)
    fmt.Println(strconv.Unquote(quoted))

}
