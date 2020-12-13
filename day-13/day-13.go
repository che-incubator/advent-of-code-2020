package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	earliestTimestamp, potentialNumbers := parseInput("./input", false)
	processBusSchedule(earliestTimestamp, potentialNumbers)
	_, potentialNumbers = parseInput("./input", true)
	findEarliestTimestamp(potentialNumbers)
}

func parseInput(inputFile string, includeX bool) (int, []int) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lineNum := 0
	var earliestTimestamp int
	potentialNumbers := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if lineNum == 0 {
			earliestTimestamp, _ = strconv.Atoi(line)
		} else {
			splitInput := strings.Split(line, ",")
			for _, v := range splitInput {
				if v != "x" {
					newNum, _ := strconv.Atoi(v)
					potentialNumbers = append(potentialNumbers, newNum)
				} else if includeX {
					potentialNumbers = append(potentialNumbers, 1)
				}
			}
		}
		lineNum += 1
	}
	return earliestTimestamp, potentialNumbers
}

func processBusSchedule(earliestTimestamp int, potentialNumbers []int) {
	closest := math.MaxInt64
	closestID := math.MaxInt64
	for _, potentialNum := range potentialNumbers {
		potentialClosest := findClosestMultiple(earliestTimestamp, potentialNum)
		if potentialClosest < closest {
			closest = potentialClosest
			closestID = potentialNum
		}
	}
	fmt.Printf("Earliest bus ID multipled by the number of minutes is: %d\n", closestID*(closest-earliestTimestamp))
}

func findEarliestTimestamp(potentialNumbers []int) {
	timestamp := 1

	for {
		timeSkip := 1
		matchFound := true

		for offset := 0; offset < len(potentialNumbers); offset++ {
			// Check if the potential number is evenly divisible
			if (timestamp+offset)%potentialNumbers[offset] != 0 {
				matchFound = false
				break
			}
			timeSkip *= potentialNumbers[offset]
		}

		if matchFound {
			fmt.Printf("The earliest timestamp is: %d\n", timestamp)
			return
		}

		timestamp += timeSkip
	}
}

func findClosestMultiple(timestamp int, num int) int {
	return int(math.Ceil(float64(timestamp)/float64(num)) * float64(num))
}
