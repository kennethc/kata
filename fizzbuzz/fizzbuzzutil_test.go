package fizzbuzzutil

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
			t.Errorf("IsFizz(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
