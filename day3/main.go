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

	s, err := logic.Split(f)
	if err != nil {
		return
	}

	// fmt.Println("REAL", s)

	a, err := logic.AssignNum(s)
	if err != nil {
		return
	}

	fmt.Println(a)

	// b, _ := logic.Split([]byte(
	// 	`vJrwpWtwJgWrhcsFMMfFFhFp
	// jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
	// PmmdzqPrVvPwwTWBwg
	// wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
	// ttgJtRGJQctTZtZT
	// CrZsJsPPZsGzwwsLwLmpwMDw`))

	// fmt.Println("FAKE", b)

	// logic.CommonType(b)

}
