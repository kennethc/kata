package main

import "testing"

type testmap []struct {
	in   int
	want string
}

var (
	cases = testmap{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{5, "Buzz"},
		{6, "Fizz"},
		{10, "Buzz"},
		{15, "FizzBuzz"},
		{30, "FizzBuzz"},
	}
)

func TestCheck(t *testing.T) {
	for _, c := range cases {
		got := check(c.in)
		if got != c.want {
			t.Errorf("check(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func BenchmarkCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		check(cases[0].in)
	}
}
