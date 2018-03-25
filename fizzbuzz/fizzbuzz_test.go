package main

import (
	"reflect"
	"testing"
)

func TestIntegers(t *testing.T) {
	cases := []struct {
		min   int
		max   int
		slice []int
	}{
		{1, 3, []int{1, 2, 3}},
	}

	for _, c := range cases {
		got := integers(c.min, c.max)
		if !reflect.DeepEqual(got, c.slice) {
			t.Errorf("integers(%q, %q) == %v, want %v", c.min, c.max, got, c.slice)
		}
	}
}
