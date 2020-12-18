package main

import (
	"fmt"
)

func main() {
	input := []int{8, 11, 0, 19, 1, 2}

	// Part 1
	game(input, 2020)

	// Part 2
	game(input, 30000000)
}

func game(initialArr []int, totalTurns int) {
	speakMap := make(map[int]int)
	initial := initialArr[0]
	var previousNum int
	totalNumsSpoken := 0
	for totalNumsSpoken < totalTurns {
		previousNum = initial

		if totalNumsSpoken < len(initialArr) {
			initial = initialArr[totalNumsSpoken]
		} else if _, ok := speakMap[initial]; ok {
			initial = totalNumsSpoken - speakMap[initial] - 1
		} else {
			initial = 0
		}

		speakMap[previousNum] = totalNumsSpoken - 1
		totalNumsSpoken += 1
	}

	fmt.Printf("For %d turns the last number spoken is: %d\n", totalTurns, initial)
}
