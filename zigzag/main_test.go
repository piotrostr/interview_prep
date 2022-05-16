package zigzag

import (
	"testing"
)

func Test1(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 3
	got := Convert(s, numRows)
	want := "PAHNAPLSIIGYIR"
	if got != want {
		t.Errorf("%s != %s", want, got)
	}
}

func Test2(t *testing.T) {
}
