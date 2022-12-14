package logic

import (
	"fmt"
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

	// fmt.Println([][]string{bag1, bag2})

	return [][]string{bag1, bag2}, nil
}

func GetValue(str string) int {
	for i := range l {
		if l[i] == str {
			return i + 1
		}
	}
	return 0
}

func AssignNum(str [][]string) (map[string]int, error) {
	ln := map[string]int{}
	t := 0
	ignore := []string{}
	list := []string{}
	newMap := map[string]bool{}

	// for i := range l {
	// 	ln[l[i]] = 0
	// }

	fmt.Println("+++++++++++++")

	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str[i]); j++ {
			for k := 0; k < len(str[i][j]); k++ {
				ln[string(str[i][j][k])] += GetValue(string(str[i][j][k]))

				if str[0][j][k] == str[1][j][k] {
					// t += GetValue(string(str[0][j][k]))
					ignore = append(ignore, string(str[0][j][k]))
				}

				// fmt.Println(i, j, k, str[i][j][k], string(str[i][j][k]))

				// fmt.Printf("\nan: %v\n", ln)

			}
		}
	}

	for i := range ignore {
		if _, value := newMap[ignore[i]]; !value {
			newMap[ignore[i]] = true
			list = append(list, ignore[i])
		}
	}

	for _, v := range list {
		fmt.Println("V", v)
		t += GetValue(v)
	}

	fmt.Println("TOTAL ", t)

	// fmt.Println("+++++++++++++")

	// fmt.Println(ln)

	return ln, nil
}
