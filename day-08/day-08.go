package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lineNumToInstruction := readInput("./input")

	// Part 1 Testing the initial branch that we know won't terminate
	_, accumulatorP1 := testBranch(lineNumToInstruction)
	fmt.Printf("The accumulator for part 1 is: %d\n", accumulatorP1)

	// Test all the variant branches for part 2 (switching exactly one nop with jmp and vice versa) and see if they terminate
	currLineNumVariant := 0
	var accumulatorP2 int
	var terminated bool
	for currLineNumVariant < len(lineNumToInstruction) {
		variant, variantMap := nextVariant(lineNumToInstruction, currLineNumVariant)
		currLineNumVariant += 1
		if !variant {
			continue
		}
		terminated, accumulatorP2 = testBranch(variantMap)
		if terminated {
			break
		}
	}

	fmt.Printf("The accumulator for the branch that terminates is: %d\n", accumulatorP2)
}

func readInput(input string) map[int]string {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineNumToInstruction := make(map[int]string)
	currLineNum := 0
	for scanner.Scan() {
		text := scanner.Text()
		lineNumToInstruction[currLineNum] = text
		currLineNum += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lineNumToInstruction
}

// If the current instruction at currLineNumVariant is no-op then switch to jmp. If its jmp then switch to no-op
// If the current line contains jmp or nop then it will return true with nop/jmp switched
// If the current line doesn't contain jmp or nop then return false
func nextVariant(lineNumToInstruction map[int]string, currLineNumVariant int) (bool, map[int]string) {
	// Copy the map so we don't change the referenced version
	duplicatedMap := make(map[int]string)
	for k,v := range lineNumToInstruction {
		duplicatedMap[k] = v
	}

	val, ok := duplicatedMap[currLineNumVariant]
	if ok {
		// If the current line contains jmp or nop then switch the instruction
		if strings.Contains(val, "jmp") {
			duplicatedMap[currLineNumVariant] = strings.Replace(val, "jmp", "nop", 1)
			return true, duplicatedMap
		} else if strings.Contains(val,"nop") {
			duplicatedMap[currLineNumVariant] = strings.Replace(val, "nop", "jmp", 1)
			return true, duplicatedMap
		} else {
			return false, duplicatedMap
		}
	}
	return false, duplicatedMap
}

/*
 * Test if the branch terminates. If it does then return the accumulated value
 */
func testBranch(lineNumToInstruction map[int]string) (bool, int){
	accumulator := 0
	previousUsedLineNumsCache := make(map[int]int)
	nextLineNum := 0
	for hasNextInstruction(previousUsedLineNumsCache, nextLineNum) {
		nextLineText := lineNumToInstruction[nextLineNum]
		previousLineNum := nextLineNum
		// This means that the program actually terminated
		if nextLineNum == len(lineNumToInstruction) {
			return true, accumulator
		}
		nextLineNum, accumulator = getNextLine(nextLineText, nextLineNum, accumulator)
		previousUsedLineNumsCache[previousLineNum] = 1
	}

	return false, accumulator
}

/*
 * Parse a line of instructions. Returns the operator (jmp, nop, acc) and the number
 */
func parseInstruction(text string) (string, int) {
	splitText := strings.Split(text, " ")
	op := splitText[0]
	add := splitText[1]
	possibleNextLine, _ := strconv.Atoi(add)
	return op, possibleNextLine
}

/*
 * getNextLine takes in a line, parses it, then returns the next line instruction that should take place
 */
func getNextLine(text string, currLineNumber int, accumulator int) (int, int){
	op, num := parseInstruction(text)
	switch op {
	case "nop":
		return currLineNumber + 1, accumulator
	case "acc":
		return currLineNumber + 1, accumulator + num
	case "jmp":
		return currLineNumber + num, accumulator
	}
	return -1, -1
}

/*
 * hasNextInstruction returns true if the instruction hasn't been executed before, false otherwise
 */
func hasNextInstruction(previousUsedLineNums map[int]int, nextLineNum int) bool {
	_, ok := previousUsedLineNums[nextLineNum]
	return !ok
}
