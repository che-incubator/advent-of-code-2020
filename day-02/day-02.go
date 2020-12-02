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
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	validPasswordCount := 0
	validPasswordCountP2 := 0
	for scanner.Scan() {
		text := scanner.Text()
		splitInputs := strings.Split(text, " ")

		// gather the lower and upper limits
		rangeInputs := strings.Split(splitInputs[0], "-")
		lower, err := strconv.Atoi(rangeInputs[0])
		if err != nil {
			log.Fatal(err)
		}
		upper, err := strconv.Atoi(rangeInputs[1])
		if err != nil {
			log.Fatal(err)
		}

		// letter with : removed
		letter := strings.Replace(splitInputs[1], ":", "", 1)

		word := splitInputs[2]

		// Check if they are valid
		if isPasswordValidP1(lower, upper, letter, word) {
			validPasswordCount +=1
		}
		if isPasswordValidP2(lower, upper, letter, word) {
			validPasswordCountP2 += 1
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("The Valid number of passwords for part 1 are: " + strconv.Itoa(validPasswordCount))
	fmt.Println("The Valid number of passwords for part 2 are: " + strconv.Itoa(validPasswordCountP2))
}

func isPasswordValidP1(lowNum int, upperNum int, letter string, word string) bool {
	letterCount := 0
	for _, char := range word {
		ch := string(char)
		if ch == letter {
			letterCount += 1
		}
	}
	return letterCount >= lowNum && letterCount <= upperNum
}

func isPasswordValidP2(lowNum int, upperNum int, letter string, word string) bool {
	lowerLetter := string(word[lowNum-1])
	upperLetter := string(word[upperNum-1])
	return (lowerLetter == letter && upperLetter != letter) || (lowerLetter != letter && upperLetter == letter)
}
