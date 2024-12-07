package main

import "testing"

func TestDec5(t *testing.T) {
	res1, res2 := dec5("../inputs/dec5_test.txt")
	if res1 != 143 {
		t.Errorf("Expected 143, got %d", res1)
	}
	if res2 != 123 {
		t.Errorf("Expected 123, got %d", res2)
	}
}

func TestMove(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	expected := []int{1, 3, 4, 2, 5}
	res := move(slice, 1, 3)
	for i, el := range res {
		if el != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], el)
		}
	}

	slice = []int{1, 2, 3, 4, 5}
	expected = []int{1, 2, 5, 3, 4}
	res = move(slice, 4, 2)
	for i, el := range res {
		if el != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], el)
		}
	}

}
