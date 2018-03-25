package main

import (
	"fmt"
	"github.com/kennethc/kata/fizzbuzz/util"
)

func main() {
	a := integers(1, 100)
	for i := 0; i < len(a); i++ {
		if util.IsFizz(a[i]) {
			fmt.Printf("Fizz")
		}
		if util.IsBuzz(a[i]) {
			fmt.Printf("Buzz")
		}
		if !util.IsFizz(a[i]) && !util.IsBuzz(a[i]) {
			fmt.Printf("%v", a[i])
		}
		fmt.Printf("\n")
	}
}

func integers(min, max int) []int {
	var r = make([]int, max-min+1)
	for i := range r {
		r[i] = i + min
	}
	return r
}
