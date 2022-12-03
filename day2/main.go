package main

import (
	"os"

	"github.com/hnucamendi/adventofcode/day2/logic"
)

// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.
// A for Rock, B for Paper, and C for Scissors.
// X for Rock, Y for Paper, and Z for Scissors.

// shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
// plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

func main() {
	f, err := os.ReadFile("./input.txt")
	if err != nil {
		return
	}

	logic.Play(f)

}
