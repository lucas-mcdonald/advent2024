package main

import (
	"fmt"
	"os"
	"strings"
)

type Dir struct {
	i int
	j int
}

func checkIsMAS(grid [][]string, i int, j int, dirs [2]Dir) bool {
	if (grid[i+dirs[0].i][j+dirs[0].j] == "M" && grid[i+dirs[1].i][j+dirs[1].j] == "S") ||
		(grid[i+dirs[0].i][j+dirs[0].j] == "S" && grid[i+dirs[1].i][j+dirs[1].j] == "M") {
		return true
	}
	return false
}

func checkIsXMas(grid [][]string, i int, j int) bool {
	if (grid[i][j] != "A" || i-1 < 0 || j-1 < 0) || (i+1 >= len(grid) || j+1 >= len(grid[i])) {
		// Early return if X-MAS can't fit
		return false
	}
	// Directions for each diagonal
	leftDirs := [2]Dir{
		{1, 1},
		{-1, -1},
	}
	rightDirs := [2]Dir{
		{1, -1},
		{-1, 1},
	}
	if checkIsMAS(grid, i, j, leftDirs) && checkIsMAS(grid, i, j, rightDirs) {
		return true
	}
	return false
}

func isNextCharValid(current string, next string) bool {
	switch current {
	case "X":
		if next == "M" {
			return true
		}
	case "M":
		if next == "A" {
			return true
		}
	case "A":
		if next == "S" {
			return true
		}
	default:
		return false
	}
	return false
}

func checkNeighbours(grid [][]string, i int, j int) int {
	total := 0
	dirs := [8]Dir{
		{0, 1},
		{1, 0},
		{1, 1},
		{0, -1},
		{-1, 0},
		{-1, -1},
		{-1, 1},
		{1, -1},
	}
	current := grid[i][j]

	if current == "X" {
		for _, dir := range dirs {
			d := 1
			var nextI, nextJ int
			nextI = i + (d * dir.i)
			nextJ = j + (d * dir.j)
			for nextI < len(grid) && nextJ < len(grid[i]) && nextI >= 0 && nextJ >= 0 {
				next := grid[nextI][nextJ]
				if isNextCharValid(current, next) {
					d += 1
					current = next
					nextI = i + (d * dir.i)
					nextJ = j + (d * dir.j)
				} else {
					break
				}
				if d == 4 {
					total += 1
					break
				}
			}
			current = grid[i][j]
		}
	}
	return total

}

func dec4(fileName string) (int, int) {
	input, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(string(input), "\n")
	var grid [][]string
	for _, line := range lines {
		row := strings.Split(string(line), "")
		grid = append(grid, row)
	}

	var p1Total int
	var p2Total int
	for i := range grid {
		for j := range grid[i] {
			p1Total += checkNeighbours(grid, i, j)
			if checkIsXMas(grid, i, j) {
				p2Total += 1
			}
		}
	}
	return p1Total, p2Total

}

func main() {
	p1Total, p2Total := dec4("../inputs/dec4.txt")
	fmt.Println("Part1: ", p1Total)
	fmt.Println("Part2: ", p2Total)
}

// [M M M S X X M A S M]
// [M S A M X M S M S A]
// [A M X S X M A A M M]
// [M S A M A S M S M X]
// [X M A S A M X A M M]
// [X X A M M X X A M A]
// [S M S M S A S X S S]
// [S A X A M A S A A A]
// [M A M M M X M M M M]
// [M X M X A X M A S X]
