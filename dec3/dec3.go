package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func get_product_from_strings(a string, b string) (product int) {
	int1, err1 := strconv.Atoi(a)
	int2, err2 := strconv.Atoi(b)
	if err1 != nil || err2 != nil {
		fmt.Println("Error: ", err1, err2)
		os.Exit(1)
	}
	return int1 * int2
}

func get_operands_from_mul(mul string) (a, b string) {
	ints := regexp.MustCompile(`\d+`).FindAllString(mul, -1)
	if len(ints) != 2 {
		fmt.Println("Error: ", ints)
		os.Exit(1)
	}
	return ints[0], ints[1]
}

func main() {
	file := "../inputs/dec3.txt"
	// Get all matches for mul\(\d\,\d\) in file
	// and print them
	input, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(string(input), -1)
	var sum int
	for _, match := range matches {
		ints := regexp.MustCompile(`\d+`).FindAllString(match, -1)
		if len(ints) != 2 {
			fmt.Println("Error: ", ints)
			os.Exit(1)
		}
		a, b := get_operands_from_mul(match)
		sum += get_product_from_strings(a, b)
	}
	fmt.Println("Part1: ", sum)

	// Part 2
	match_indices := re.FindAllStringIndex(string(input), -1)

	do_re := regexp.MustCompile(`do\(\)`)
	dont_re := regexp.MustCompile(`don\'t\(\)`)
	do_indices := do_re.FindAllStringIndex(string(input), -1)
	dont_indices := dont_re.FindAllStringIndex(string(input), -1)

	var en bool = true
	var sum2 int
	var do_idx, dont_idx, match_idx int

	for i := range input {
		if do_idx < len(do_indices) && i == do_indices[do_idx][0] {
			en = true
			do_idx += 1
			continue
		}
		if dont_idx < len(dont_indices) && i == dont_indices[dont_idx][0] {
			en = false
			dont_idx += 1
			continue
		}
		if match_idx < len(match_indices) && i == match_indices[match_idx][0] {
			if en {
				match := match_indices[match_idx]
				a, b := get_operands_from_mul(string(input[match[0]:match[1]]))
				sum2 += get_product_from_strings(a, b)
			}
			match_idx += 1
		}
	}
	fmt.Println("Part2", sum2)
}
