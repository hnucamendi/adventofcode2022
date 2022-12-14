package logic

import (
	"strings"
)

var l = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func Split(s []byte) ([][]string, error) {
	var (
		bag1 []string
		bag2 []string
	)

	nt := strings.ReplaceAll(string(s), "\t", "")
	f := strings.Split(nt, "\n")

	for _, t := range f {
		bag1 = append(bag1, t[:len(t)/2])
		bag2 = append(bag2, t[len(t)/2:])
	}

	return [][]string{bag1, bag2}, nil
}

func AssignNum(str [][]string) (int, error) {
	t := 0

	for i := 0; i < len(str[0]); i++ {
		for j := range str[0][i] {
			if strings.Contains(str[1][i], string(str[0][i][j])) {
				t += GetValue(string(str[0][i][j]))
				break
			}
		}
	}
	return t, nil
}

func GetValue(str string) int {
	for i := range l {
		if l[i] == str {
			return i + 1
		}
	}
	return 0
}
