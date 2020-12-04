package main

import (
	readAOC "../utils"
	"fmt"
	"math"
	"strconv"
)

func main() {
	lines := readAOC.ReadInput("input.txt")
	var inputIntArr []int
	for _, num := range lines {
		intVal, _ := strconv.Atoi(num)
		inputIntArr = append(inputIntArr, intVal)
	}
	fmt.Printf("Part 1: %d\n", Part1(inputIntArr))
	fmt.Printf("Part 2: %d\n", Part2(inputIntArr))
}

func Part1(inputArr []int) int {
	for i, v := range inputArr {
		for j, v2 := range inputArr {
			if i == j {
				continue
			}
			if v + v2 == 2020 {
				return v * v2
			}
		}
	}
	return 0
}

func Part2(inputArr []int) int {
	prunedInput := getPrunedInput(inputArr)
	for i, v := range prunedInput {
		for j, v2 := range prunedInput {
			if v + v2 > 2020 || i == j {
				continue
			}
			for k, v3 := range prunedInput {
				if j == k || j == i {
					continue
				}
				if v + v2+ v3 == 2020 {
					return v * v2 * v3
				}
			}
		}
	}
	return 0
}

// Ignore any number + i that's > 2020
func getPrunedInput(inputArr []int) []int {
	var prunedInput []int
	maxInt := math.MaxInt64
	for _, val := range inputArr {
		if val < maxInt {
			maxInt = val
		}
	}
	for _, val := range inputArr {
		if val + maxInt < 2020 {
			prunedInput = append(prunedInput, val)
		}
	}
	return prunedInput
}
