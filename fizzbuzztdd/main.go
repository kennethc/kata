package main

import (
    "strconv"
    "fmt"
)

func main() {
    for i := 1; i <= 100; i++ {
        fmt.Println(check(i))
    }
}

func check(x int) string {
    var r string
    if x%3 == 0 {
        r += "Fizz"
    }
    if x%5 == 0 {
        r += "Buzz"
    }
    if r == "" {
        r = strconv.Itoa(x)
    }
    return r
}
