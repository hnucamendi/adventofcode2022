package logic

import (
	"fmt"
	"strings"
	"time"
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

func removeDuplicate(dup []string) []string {
	newMap := map[string]bool{}
	list := []string{}

	for i := range dup {
		if _, value := newMap[dup[i]]; !value {
			newMap[dup[i]] = true
			list = append(list, dup[i])
		}
	}
	return list
}

func AssignNum(str [][]string) (map[string]int, error) {
	ln := map[string]int{}
	ignore := []string{}
	t := 0

	fmt.Println("+++++++++++++")

	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str[i]); j++ {
			for k := 0; k < len(str[i][j]); k++ {
				ln[string(str[i][j][k])] += GetValue(string(str[i][j][k]))
				fmt.Printf("ln:%v str:%v value:%v\n", ln[string(str[i][j][k])], string(str[i][j][k]), GetValue(string(str[i][j][k])))
				time.Sleep(time.Millisecond * 350)

				if str[0][j][k] == str[1][j][k] {
					// t += GetValue(string(str[0][j][k]))
					ignore = append(ignore, string(str[0][j][k]))
				}
			}
		}
	}

	list := removeDuplicate(ignore)
	for _, v := range list {
		// fmt.Printf("%s %v + %v = %v\n", v, t, GetValue(v), t+GetValue(v))
		t += GetValue(v)
		// time.Sleep(time.Millisecond * 350)
	}

	fmt.Println("TOTAL ", t)

	// fmt.Println("+++++++++++++")

	// fmt.Println(ln)

	return ln, nil
}
