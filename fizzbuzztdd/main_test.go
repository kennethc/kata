package main

import "testing"

func TestCheck(t *testing.T) {
    cases := []struct {
        in int
        want string
    }{
        {1, "1"},
        {2, "2"},
        {3, "Fizz"},
        {5, "Buzz"},
        {6, "Fizz"},
        {10, "Buzz"},
        {15, "FizzBuzz"},
        {30, "FizzBuzz"},
    }

    for _, c := range cases {
        got := check(c.in)
        if got != c.want {
            t.Errorf("check(%v) == %v, want %v", c.in, got, c.want)
        }
    }
}
