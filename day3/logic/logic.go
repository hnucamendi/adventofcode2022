package logic

import (
	"fmt"
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

	// fmt.Println([][]string{bag1, bag2})

	return [][]string{bag1, bag2}, nil
}

func AssignNum(str [][]string) (map[string]int, error) {
	newMap := make(map[string]bool)
	similar := []string{}
	t := 0
	l := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	an := map[string]int{}
	ln := map[string]int{}

	for i := 0; i < len(l); i++ {
		an[l[i]] = i + 1
		// ln[l[i]] = i + 1
	}

	// fmt.Println(an)
	fmt.Println("+++++++++++++")

	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str[i]); j++ {
			for k := 0; k < len(str[i][j]); k++ {
				ln[string(str[i][j][k])] += an[string(str[i][j][k])]

				if str[0][j][k] == str[1][j][k] {
					// fmt.Println(string(str[0][j][k]))
					similar = append(similar, string(str[0][j][k]))
				}

				// fmt.Println(i, j, k, str[i][j][k], string(str[i][j][k]))

				// fmt.Printf("\nan: %v\n", ln)

			}
		}
	}

	list := []string{}
	for i := range similar {
		if _, value := newMap[similar[i]]; !value {
			newMap[similar[i]] = true
			list = append(list, similar[i])
		}
	}

	for i := range list {
		t += ln[list[i]]
	}

	fmt.Print("TOTAL ", list, t)

	// fmt.Println("+++++++++++++")

	// fmt.Println(ln)

	return ln, nil
}
