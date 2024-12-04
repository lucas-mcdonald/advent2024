package main

import "testing"

func TestDec4(t *testing.T) {
	res1, res2 := dec4("inputs/dec4_test.txt")
	if res1 != 18 {
		t.Errorf("Expected 18, got %d", res1)
	}
	if res2 != 9 {
		t.Errorf("Expected 9, got %d", res2)
	}
}
