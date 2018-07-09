package main

import "testing"

func TestSum(t *testing.T) {
	var total = sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}


func sum(x int, y int) int {
	return x + y
}
