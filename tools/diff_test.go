package tools

import (
	"testing"
)

func TestDiffString(t *testing.T) {
	n := []string{"a", "b"}
	o := []string{"a", "c"}
	nn, oo := DiffString(n, o)
	t.Log(nn, oo)
}

func TestDiffInt(t *testing.T) {
	n := []int{1, 2}
	o := []int{1, 3}
	nn, oo := DiffInt(n, o)
	t.Log(nn, oo)
}
