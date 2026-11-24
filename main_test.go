package main

import "testing"

func TestKthPermutation(t *testing.T) {
	cases := []struct {
		n, k int
		want string
	}{
		{3, 3, "213"},
		{4, 9, "2314"},
		{3, 1, "123"},
		{3, 6, "321"},
		{1, 1, "1"},
		{4, 17, "3412"},
		{3, 7, ""}, // k out of range
		{0, 1, ""}, // n out of range
	}
	for _, c := range cases {
		if got := KthPermutation(c.n, c.k); got != c.want {
			t.Errorf("KthPermutation(%d,%d) = %q, want %q", c.n, c.k, got, c.want)
		}
	}
}
