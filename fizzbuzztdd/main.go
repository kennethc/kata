package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(check(i))
	}
}

func check(x int) string {
	var r string
	if multiple(x, 3) {
		r += "Fizz"
	}
	if multiple(x, 5) {
		r += "Buzz"
	}
	if r == "" {
		r = strconv.Itoa(x)
	}
	return r
}

func multiple(x int, m int) bool {
	return x%m == 0
}
