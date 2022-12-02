package main

import (
	"github.com/DOI/adventofcode2022/day1/challenge1/challengeOne/elfs"
)

func main() {
	// f, err := os.ReadFile("./input.txt")
	// if err != nil {
	// 	return
	// }

	// elfs.ElfParser(f)

	elfs.ElfParser([]byte(`
	123
	456

	321
	432
	`))
}
