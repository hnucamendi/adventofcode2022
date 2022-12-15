package logic

import (
	"strings"
)

var l = []string{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"}

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
	for i := range l {
		return strings.Index(l[i], str) + 1
	}
	return 0
}
