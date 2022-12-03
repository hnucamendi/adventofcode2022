package logic

import (
	"fmt"
	"strings"
)

type PlayOption string

var (
	rock     int = 1
	paper    int = 2
	scissors int = 3
	win      int = 6
	draw     int = 3
	lose     int = 0
)

func Play(p []byte) (int, int, error) {
	var comp int
	var usr int

	f := strings.Split(string(p), "\n")

	for i := 0; i < len(f); i++ {
		f[i] = strings.TrimSpace(f[i])

		switch f[i] {
		case "A X":
			comp += rock + draw
			usr += rock + draw
		case "A Y":
			comp += rock + lose
			usr += paper + win
		case "A Z":
			comp += rock + win
			usr += scissors + lose
		case "B X":
			comp += paper + win
			usr += rock + lose
		case "B Y":
			comp += paper + draw
			usr += paper + draw
		case "B Z":
			comp += paper + lose
			usr += scissors + win
		case "C X":
			comp += scissors + lose
			usr += rock + win
		case "C Y":
			comp += scissors + win
			usr += paper + lose
		case "C Z":
			comp += scissors + draw
			usr += scissors + draw
		}
	}

	fmt.Println("COMP", comp)
	fmt.Println("USR", usr)

	return comp, usr, nil
}
