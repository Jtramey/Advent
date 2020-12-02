package main

import (
	readAOC "advent/utils"
	"fmt"
	"strconv"
	"strings"
)

type Password struct {
	LowerBound int
	UpperBound int
	Letter     string
	Password   string
}

func main() {
	lines := readAOC.ReadInput("input.txt")
	part1Count := 0
	part2Count := 0
	for _, num := range lines {
		p := Password{}
		// Split input into useful things
		halves := strings.Split(num, ": ")
		p.Password = strings.Trim(" ", halves[1])
		rules := strings.Split(halves[0], " ")
		p.Letter = rules[1]
		bounds := strings.Split(rules[0], "-")
		p.LowerBound, _ = strconv.Atoi(bounds[0])
		p.UpperBound, _ = strconv.Atoi(bounds[1])
		if p.isValid() {
			part1Count++
		}
		if p.isValidPart2() {
			part2Count++
		}
	}
	fmt.Printf("Part 1: %d\n", part1Count)
	fmt.Printf("Part 1: %d\n", part2Count)
}

func (p *Password) isValid() bool {
	appearances := strings.Count(p.Password, p.Letter)
	return appearances >= p.LowerBound && appearances <= p.UpperBound
}

func (p *Password) isValidPart2() bool {
	lowerBoundMatches := string(p.Password[p.LowerBound-1]) == p.Letter
	upperBoundMatches := string(p.Password[p.UpperBound-1]) == p.Letter
	return !(lowerBoundMatches && upperBoundMatches) && (lowerBoundMatches || upperBoundMatches)
}
