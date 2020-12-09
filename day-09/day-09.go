package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := readInput("./input")
	firstInvalid := firstInvalid(numbers, 25)
	log.Println(firstInvalid)

	weaknessNumber := findWeakness(numbers, firstInvalid)
	log.Println(weaknessNumber)
}

func firstInvalid(numbers []int, preamble int) int {
	for i := preamble; i < len(numbers); i = i + 1 {
		if !findSum(numbers[i-preamble:i], numbers[i]) {
			return numbers[i]
		}
	}
	return -1
}

func findWeakness(numbers []int, invalidNo int) int {
	for i, no := range numbers {
		sum := no
		if sum > invalidNo {
			continue
		}
		// log.Println("starting with sum", sum)
		for j, no2 := range numbers[i+1:] {
			sum = sum + no2
			// log.Printf("added %d, current sum is %d\n", no2, sum)
			if sum == invalidNo {
				// log.Println("found", sum, i, j)
				return sumMinMax(numbers[i:i + j + 1])
			} else if sum > invalidNo {
				// log.Println(sum, "is too much. ", invalidNo)
				break
			}
		}
	}
	return -1
}

func sumMinMax(numbers []int) int {
	min, max := numbers[0], numbers[0]

	for _, no := range numbers[1:] {
		if no < min {
			min = no
		}
		if no > max {
			max = no
		}
	}
	return min + max
}

func findSum(numbers []int, sum int) bool {
	for i, no := range numbers {
		for _, no2 := range numbers[i:] {
			if no + no2 == sum {
				return true
			}
		}
	}
	return false
}

func readInput(input string) []int {
	numbers := make([]int, 0)

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// log.Println(line)

		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return numbers
}
