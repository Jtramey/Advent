package main

import (
	readAOC "../utils"
	"fmt"
)

func main() {
	lines := readAOC.ReadInput("")
	grid := asRunes(lines)

	slopes := []slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	var results []int
	for _, slope := range slopes {
		x, y, treeCount := 0, 0, 0
		for y < len(grid) {
			if grid[y][x] == '#' {
				treeCount++
			}
			x = (x + slope.Dx) % len(grid[0])
			y = y + slope.Dy
		}
		fmt.Println("# of trees met =", treeCount, fmt.Sprintf("(slope %v)", slope))
		results = append(results, treeCount)
	}

	fmt.Println("Product of results =", reduce(results, func(acc, next int) int { return acc * next }))
}

func asRunes(lines []string) [][]rune {
	var a [][]rune
	for _, line := range lines {
		var b []rune
		for _, letter := range line {
			b = append(b, letter)
		}
		a = append(a, b)
	}
	return a
}

type slope struct {
	Dx int
	Dy int
}

func reduce(nums []int, f func(acc, next int) int) int {
	o := nums[0]
	for _, i := range nums[1:] {
		o = f(o, i)
	}
	return o
}
