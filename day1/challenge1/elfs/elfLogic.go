package elfs

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Topthree(arr [][]int) (int, error) {
	var sum int = 0

	for i := 0; i < 3; i++ {
		sum += arr[i][0]
	}

	return sum, nil
}

func SortArr(arr [][]int) ([][]int, error) {

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] > arr[j][0]
	})

	return nil, nil
}

// converts a string of values into multidimensional array of elfs
func ElfParser(bte []byte) ([][]int, error) {

	var l [][]int
	var count int = 0

	o := strings.ReplaceAll(string(bte), "\t", "")
	o = strings.ReplaceAll(o, "\n", " ")
	o = strings.TrimSpace(o)
	f := strings.Split(o, " ")

	for i := 0; i < len(f); i += 1 {
		if f[i] == "" {
			l = append(l, []int{count})
			count = 0
			i += 1
		}

		n, err := strconv.Atoi(f[i])
		CheckErr(err)

		count += n

		if i == len(f)-1 {
			l = append(l, []int{count})
			count = 0
		}
	}
	return l, nil
}

func CheckErr(err error) ([][]int, error) {
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	return nil, nil
}
