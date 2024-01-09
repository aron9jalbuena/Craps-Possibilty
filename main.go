package main

import (
	"fmt"
	"math/rand"
	"sort"
)

const (
	min = 1
	max = 7

	num_of_rolls = 100
)

func roll() int {
	return rand.Intn(max-min) + min // generate a random number between min and max
}

func main() {
	chances := map[int]int{}
	for i := 0; i <= num_of_rolls; i++ {
		dice := roll() + roll()
		if _, found := chances[dice]; found {
			chances[dice]++
		} else {
			chances[dice] = 1
		}
	}

	// Sort numbers in an array
	numbers := []int{}
	for num := range chances {
		numbers = append(numbers, num)
	}
	sort.Ints(numbers)

	// Print out assortedly
	for _, key := range numbers {
		fmt.Printf("%d:%d\n", key, chances[key])
	}
}
