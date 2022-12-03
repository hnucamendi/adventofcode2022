package main

import (
	"fmt"
	"os"

	"github.com/hnucamendi/adventofcode/day2/logic"
)

func main() {
	f, err := os.ReadFile("./input.txt")
	if err != nil {
		return
	}

	fmt.Println("Part One")
	logic.Play(f)
	fmt.Println("----------")
	fmt.Println("Part Two")
	logic.PlayR(f)
}
