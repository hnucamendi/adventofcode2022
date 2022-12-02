package elfs

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Elfs struct {
	Elf []string
}

// Sums up the values in an array
// In this instance used to determine the total calories each elf has
func CalorieCounter(calories []int) (int, error) {

	// fmt.Println("Before", calories)

	sort.SliceStable(calories, func(i, j int) bool {
		return true
	})

	// fmt.Println("After", calories)

	// arr.sort(function (a, b) {
	// 	return a - b;
	// });

	sum := 0
	for _, c := range calories {
		sum += c
	}

	return sum, nil
}

// converts a string of values into multidimensional array of elfs
func ElfParser(bte []byte) ([]interface{}, error) {
	var l []interface{}

	o := strings.Split(string(bte), "\n\n")

	fmt.Println("spaced", o[1])
	fmt.Println("spaced", []byte(o[1]))

	for i := 0; i < len(o); i++ {
		n := strings.Split(o[i], "\n")
		for j := 0; j < len(n); j++ {
			ls := strings.ReplaceAll(n[j], "\t", "")
			if ls == "" {
				continue
			}
			s, err := strconv.Atoi(ls)
			if err != nil {
				fmt.Println(err)
			}
			l = append(l, s)
		}
		fmt.Println(o[i])
		fmt.Println([]byte(o[i]))
	}

	fmt.Println(l)

	return nil, nil
	for i := 0; i < len(o); i++ {
		s := strings.Split(o[i], "\n\t\t\t\t\t")
		l = append(l, s)
	}
	fmt.Println("HERE", l)

	f, _ := json.Marshal(l)
	fmt.Println("here", string(f))

	return l, nil
}
