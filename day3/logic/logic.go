package logic

import (
	"fmt"
	"strings"
)

func Split(s []byte) ([]string, []string, error) {
	var bag1 []string
	var bag2 []string

	nt := strings.ReplaceAll(string(s), "\t", "")
	f := strings.Split(nt, "\n")

	for _, t := range f {
		bag1 = append(bag1, t[:len(t)/2])
		bag2 = append(bag2, t[len(t)/2:])
	}

	return bag1, bag2, nil
}

func CommonType(bag1, bag2 []string) (string, error) {
	var count int

	for i := 0; i < len(bag1); i++ {
		for j := 0; j < len(bag1[i]); j++ {
			fmt.Println(i, bag1, j, bag1[i][j])
		}
	}
	fmt.Println(count)

	return "", nil
}
