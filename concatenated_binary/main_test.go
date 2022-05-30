package main

import (
	"testing"
)

func TestConcatenatedBinary(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{1, 1},
		{3, 27},
		{12, 505379714},
	}
	for _, c := range cases {
		if got := ConcatenatedBinary(c.in); got != c.want {
			msg := "ConcatenatedBinary(%d) == %d, want %d"
			t.Errorf(msg, c.in, got, c.want)
		}
	}
}
