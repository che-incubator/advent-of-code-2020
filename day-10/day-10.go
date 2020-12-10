package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	numbers := readInput("./input")
	sort.Ints(numbers)
	numbers = append([]int{0}, numbers...)	// add beginning 0
	numbers = append(numbers, numbers[len(numbers) - 1] + 3)	// add device jolts highest + 3

	diffs := diffs(numbers)
	result := diffs[1] * diffs[3]
	log.Println("What is the number of 1-jolt differences multiplied by the number of 3-jolt differences?")
	log.Println("Answer:", result)
	log.Println("=============================")
	log.Println()

	arrangments := calcArrangments(numbers, make(map[int]int))
	log.Println("What is the total number of distinct ways you can arrange the adapters to connect the charging outlet to your device?")
	log.Println("Answer:", arrangments)
	log.Println("=============================")

}

func calcArrangments(numbers []int, cache map[int]int) int {
	arrangments := 1
	for i, _ := range numbers {
		if i == 0 || i == len(numbers) - 1 {
			continue
		}

		// log.Println(numbers)
		tmp := append([]int(nil), numbers...)
		try := append(tmp[:i], tmp[i + 1:]...)
		diffs := diffs(try)
		if diffs[4] == 0 && diffs[5] == 0 && diffs[6] == 0 {
			// log.Println(try)
			sum := sum(try[i - 1:])	// can we use sum of the slice as key to the cache? It works for my input
			if _, ok := cache[sum]; ok {
				arrangments = arrangments + cache[sum]
			} else {
				arr := calcArrangments(try[i - 1:], cache)
				arrangments = arrangments + arr
				cache[sum] = arr
			}
		}
	}
	return arrangments
}

func sum(slice []int) int {
	sum := 0
	for _, no := range slice {
		sum = sum + no
	}
	return sum
}

func diffs(numbers []int) map[int]int {
	diffs := make(map[int]int)
	for i, no := range numbers {
		if i == 0 {
			continue
		}

		diff := no - numbers[i - 1]
		diffs[diff] = diffs[diff] + 1
	}
	return diffs
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
