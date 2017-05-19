package stringutil

import (
  "testing"
  "fmt"
)

func TestTest(t *testing.T) {
	cases := []struct {
		in string
	}{
		{"Hello, world"},
		{"Hello, ??"},
		{""},
	}
	for _, c := range cases {
    fmt.Print(c.in)
	}
}