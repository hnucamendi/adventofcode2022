package main

import (
	"github.com/hnucamendi/adventofcode/day3/logic"
)

func main() {

	// f, err := os.ReadFile("./input.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// logic.Split(f)

	b1, b2, err := logic.Split([]byte(
		`vJrwpWtwJgWrhcsFMMfFFhFp
	jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
	PmmdzqPrVvPwwTWBwg
	wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
	ttgJtRGJQctTZtZT
	CrZsJsPPZsGzwwsLwLmpwMDw`))

	// fmt.Printf("Bag1 %v\nBag2 %v\n", b1, b2)

	if err != nil {
		return
	}

	logic.CommonType(b1, b2)

}
