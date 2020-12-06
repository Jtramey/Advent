package main

import (
	"../utils"
	"fmt"
	"github.com/Sirupsen/logrus"
	"sort"
	"strconv"
	"strings"
)

type seat struct {
	row int
	col int
}

func main() {
	lines := readAOC.ReadInput("")

	var ids []int

	for _, line := range lines {
		seat := readBoardingPass(line)
		ids = append(ids, seat.id())
	}

	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	fmt.Println("Part 1: ", ids[len(ids)-1])

	for idx, id := range ids {
		if ids[idx+1] != id+1 {
			fmt.Println("Part 2: ", id + 1)
			return
		}
	}
}

func (s *seat) id() int {
	return s.row*8 + s.col
}

func readBoardingPass(s string) seat {
	s = strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1").Replace(s)
	row64, err := strconv.ParseInt(s[:7], 2, 8)
	if err != nil {
		logrus.Error("RIP")
	}
	col64, err := strconv.ParseInt(s[7:], 2, 8)
	if err != nil {
		logrus.Error("RIP")
	}
	return seat{
		row: int(row64),
		col: int(col64),
	}
}
