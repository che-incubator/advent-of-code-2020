package main

import (
	"log"
	"bufio"
	"os"
	"fmt"
)

func main() {
	treeMap := readMap("./input")
	// printMap(treeMap)
	
	slope11 := treesMet(treeMap, 1, 1)
	log.Println("Slope [1,1] = ", slope11)

	slope31 := treesMet(treeMap, 3, 1)
	log.Println("Slope [3,1] = ", slope31)

	slope51 := treesMet(treeMap, 5, 1)
	log.Println("Slope [5,1] = ", slope51)
	
	slope71 := treesMet(treeMap, 7, 1)
	log.Println("Slope [7,1] = ", slope71)

	slope12 := treesMet(treeMap, 1, 2)
	log.Println("Slope [1,2] = ", slope12)

	log.Println("multiply = ", slope11 * slope31 * slope51 * slope71 * slope12)
}

func treesMet(treeMap [][]bool, slopeRight int, slopeDown int) int {
	treesMet := 0
	col := 0
	for row := 0; row < len(treeMap); row = row + slopeDown {
		if treeMap[row][col % len(treeMap[0])] {
			treesMet = treesMet + 1
		}
		col = col + slopeRight
	}
	return treesMet
}

func readMap(inputFile string) [][]bool {
	treeMap := make([][]bool, 0)

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lineNo := 0
	for scanner.Scan() {
		line := scanner.Text()

		treeMap = append(treeMap, make([]bool, len(line)))
		for i, field := range []rune(line) {
			treeMap[lineNo][i] = field == '#'
		}
		lineNo = lineNo + 1
	}
	
	return treeMap
}

func printMap(treeMap [][]bool) {
	for _, mapLine := range treeMap {
		for _, mapField := range mapLine {
			if mapField {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}