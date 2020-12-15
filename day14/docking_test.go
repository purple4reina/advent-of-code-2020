package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	input := strings.Split(`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`, "\n")
	answer := Solve(input)
	if answer != 208 {
		t.Errorf("incorrect answer: expect=208, actual=%d", answer)
	}
}
