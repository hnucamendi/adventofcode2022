package main

func CalorieCounter(calories []int) (int, error) {

	sum := 0
	for _, c := range calories {
		sum += c
	}

	return sum, nil
}
