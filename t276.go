// 枚举
package main

const (
    Sunday = iota
    Monday
    Tuesday
    Wednesay
    Thursday
    Friday
    Saturday
)

const (
    _ = iota
    KB int64 = 1 << (10* iota)
    MB
    GB
)

const (
    A,B = iota,iota <<10
    C,D
)

func main(){
    println(KB,MB,GB)
    println(A,B,C,D)

}

