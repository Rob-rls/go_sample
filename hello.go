package main

import (
    "fmt"
    "math/cmplx"
)


func swap(x, y string) (string, string) {
    return y, x
}

func add(x, y int) int {
  return x + y
}

func split(sum int) (x, y int) {   
    x = sum * 4 / 9
    y = sum - x
    return x, y
}

var (
    ToBe bool = false
    MaxInt uint64 = 1<<64 - 1
    z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
    var i, j int = 1, 2
    k := 3
    c, python, java := true, false, "no!"
    fmt.Println(i, j, k, c, python, java)
    fmt.Println(add(42, 13))
    a, b := swap("hello", "world")
    fmt.Println(a, b)
    fmt.Println(split(17))
    fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
    fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
    fmt.Printf("Type: %T Value: %v\n", z, z)
}

