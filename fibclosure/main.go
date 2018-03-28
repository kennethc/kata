package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	p, q, r, s := 0, 0, 1, 0
	return func() int {
		s = q + r
		p = q
		q = r
		r = s
		return p
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(i, ": ", f())
	}
}
