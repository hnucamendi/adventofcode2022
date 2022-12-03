package main

import (
	"fmt"
	"os"

	"github.com/DOI/adventofcode2022/day1/elfs"
)

func main() {
	f, err := os.ReadFile("./input.txt")
	if err != nil {
		return
	}

	highestCalories, err := elfs.ElfParser(f)
	elfs.CheckErr(err)

	tp, err := elfs.Topthree(highestCalories)
	elfs.CheckErr(err)

	// Part 1
	fmt.Println(highestCalories)

	// Part 2
	fmt.Println("TOP 3:", tp)

}
