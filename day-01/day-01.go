package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	numbers := readInputNumbers("./input")
	sort.Ints(numbers)
	log.Println("found product of pair with sum 2020 =", findSumProduct2(numbers, 2020))
	log.Println()
	log.Println("found product of triplet with sum 2020 =", findSumProduct3(numbers, 2020))
}

func findSumProduct3(sortedNumbers []int, find int) int {
	for _, n1 := range sortedNumbers {
		findRemains := find - n1
		found := findSumProduct2(sortedNumbers, findRemains)
		if found != -1 {
			log.Printf("found pair for [%d]", n1)
			return n1 * found
		}
	}
	log.Printf("not found triplet with sum [%d]", find)
	return -1
}

func findSumProduct2(sortedNumbers []int, find int) int {
	for i, n1 := range sortedNumbers {
		for j := len(sortedNumbers) - 1; j > i; j-- {
			n2 := sortedNumbers[j]
			sum := n1 + n2
			if sum < find {
				break
			} else if sum == find {
				log.Printf("found sum [%d] = [%d] + [%d]", find, n1, n2)
				return n1 * n2
			}
		}
	}
	log.Printf("not found pair with sum [%d]", find)
	return -1
}

func readInputNumbers(inputFile string) []int {
	numbers := make([]int, 0)
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
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
