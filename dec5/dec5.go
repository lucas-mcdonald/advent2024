package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findIndex(arr []int, target int) int {
	for i, el := range arr {
		if el == target {
			return i
		}
	}
	return -1
}

func move(arr []int, from int, to int) []int {
	// Handle invalid indices
	if from < 0 || from >= len(arr) || to < 0 || to >= len(arr) {
		return arr
	}

	// Store the element to move
	elementToMove := arr[from]

	// Create new slice removing the 'from' element
	newArr := append(arr[:from], arr[from+1:]...)

	// Insert element at 'to' position
	return append(newArr[:to], append([]int{elementToMove}, newArr[to:]...)...)
}

func to_ints(arr []string) (ints []int) {
	for i, el := range arr {
		intEl, err := strconv.Atoi(el)
		if err != nil {
			fmt.Printf("Cannot convert string at index %d: %s\n", i, el)
		}
		ints = append(ints, intEl)
	}
	return
}

func is_ok(printing []int, rulesMap map[int][]int) bool {
	ok := true
	for pos, page := range printing {
		following := rulesMap[page]
		for _, number := range following {
			idx := findIndex(printing, number)
			if idx != -1 && idx < pos {
				ok = false
				// Fix it
			}
		}
	}
	return ok
}

func reorder(printing []int, rulesMap map[int][]int) ([]int, bool) {
	ok := true
	for pos, page := range printing {
		following := rulesMap[page]
		for _, number := range following {
			idx := findIndex(printing, number)
			if idx != -1 && idx < pos {
				printing = move(printing, pos, idx)
				pos -= 1
				ok = false
			}
		}
	}
	return printing, ok
}

func dec5(fileName string) (p1 int, p2 int) {
	input, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(string(input), "\n")

	rulesMap := make(map[int][]int)
	var printings [][]int
	is_rules := true

	for _, line := range lines {
		if line == "" {
			is_rules = false
			continue
		}
		if is_rules {
			split_line := to_ints(strings.Split(line, "|"))
			rulesMap[split_line[0]] = append(rulesMap[split_line[0]], split_line[1])
		} else {
			printings = append(printings, to_ints(strings.Split(line, ",")))
		}
	}
	for _, printing := range printings {
		ok := is_ok(printing, rulesMap)
		if ok {
			p1 += printing[len(printing)/2]
		} else {
			for !ok {
				// Fix
				printing, ok = reorder(printing, rulesMap)
			}
			p2 += printing[len(printing)/2]
		}

	}
	return
}

func main() {
	p1Total, p2Total := dec5("./inputs/dec5.txt")
	fmt.Println("Part1: ", p1Total)
	fmt.Println("Part2: ", p2Total)
}
