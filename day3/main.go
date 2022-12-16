package main

import (
	"fmt"
	"os"

	"github.com/hnucamendi/adventofcode/day3/logic"
)

func main() {
	f, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Part 1
	s, err := logic.Split(f)
	if err != nil {
		return
	}

	a, err := logic.AssignNum(s)
	if err != nil {
		return
	}
	fmt.Println("Part 1", a)

	// Part 2
	se, err := logic.Split3Ways(f)
	if err != nil {
		return
	}

	ae, err := logic.AssignNum3Ways(se)
	if err != nil {
		return
	}

	fmt.Println("Part 2", ae)

}
