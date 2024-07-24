package qns6_7

import (
	"testing"
)

func TestApocalypse(t *testing.T) {
	res := Apocalypse()
	if abs(1, res) > 0.05 {
		t.Errorf("Population should be close to 1, got = %v", res)
	}
}

type Number interface {
	int | float64
}

func abs[T Number](a, b T) T {
	if a > b {
		return a - b
	}
	return b - a
}
