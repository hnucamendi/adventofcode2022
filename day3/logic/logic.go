package logic

import (
	"strings"
)

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
	var t int = 0

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
	var l = []string{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"}
	for i := range l {
		return strings.Index(l[i], str) + 1
	}
	return 0
}

func Split3Ways(s []byte) ([][]string, error) {
	var (
		oLength int
		bag     [][]string
	)

	nt := strings.ReplaceAll(string(s), "\t", "")
	f := strings.Split(nt, "\n")

	oLength = len(f) / 3

	for i := 0; i < oLength; i++ {
		bag = append(bag, []string{f[0], f[1], f[2]})
		f = f[3:]
	}

	return bag, nil
}

func AssignNum3Ways(bag [][]string) (int, error) {
	var t int

	for _, b := range bag {
		for i := 0; i < len(b); i += 3 {
			for _, bii := range b[i] {
				if strings.Contains(b[0], string(bii)) && strings.Contains(b[1], string(bii)) && strings.Contains(b[2], string(bii)) {
					t += GetValue(string(bii))
					break
				}
			}
		}
	}
	return t, nil
}
