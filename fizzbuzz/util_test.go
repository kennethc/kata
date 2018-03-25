package util

import "testing"

func TestIsFizz(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{1, false},
		{3, true},
	}

	for _, c := range cases {
		got := IsFizz(c.in)
		if got != c.want {
			t.Errorf("IsFizz(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}

func TestIsBuzz(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{1, false},
		{5, true},
	}

	for _, c := range cases {
		got := IsBuzz(c.in)
		if got != c.want {
			t.Errorf("IsBuzz(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}
