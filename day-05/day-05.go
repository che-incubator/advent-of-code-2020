package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	readSeats("./input")
}

func readSeats(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	highestSeatID := 0
	var seatIDs []int
	for scanner.Scan() {
		line := scanner.Text()
		row := findRow(line, 0, 0, 127)
		column := findColumn(line[7:], 0,  0, 7)
		seatID := int(row * 8 + column)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
		seatIDs = append(seatIDs, seatID)
	}
	fmt.Printf("The highest seat id is: %d\n", highestSeatID)
	sort.Ints(seatIDs)
	yourSeatID := findYourSeatID(seatIDs)
	fmt.Printf("Your seat is id: %d\n", yourSeatID)
}

/**
 * findRow Gets the row by recursively looking through the letters
 * if the letter is F take the lower half
 * if the letter is B take the upper half
 * if the letter is the last letter then return lower if its F or higher if its B
 */
func findRow(letters string, currentLetterIndex int, lower float64, higher float64) float64 {
	letter := string(letters[currentLetterIndex])

	// Ending case where we find the row
	if currentLetterIndex == 6 {
		if letter == "F" {
			return lower
		} else {
			return higher
		}
	}

	if letter == "F" {
		return findRow(letters, currentLetterIndex + 1, lower, math.Floor((lower + higher) / 2))
	} else {
		// Letter is B
		return findRow(letters, currentLetterIndex + 1, math.Ceil((lower + higher)/2), higher)
	}
}

/**
 * findColumn Gets the column by recursively looking through the letters.
 * if the letter is L take the lower half
 * if the letter is R take the upper half
 * if the letter is the last letter then return lower if its L or higher if its R
 */
func findColumn(letters string, currentLetterIndex int, lower float64, higher float64) float64 {
	letter := string(letters[currentLetterIndex])

	// Ending case where we find the column
	if currentLetterIndex == 2 {
		if letter == "L" {
			return lower
		} else {
			return higher
		}
	}

	if letter == "L" {
		return findColumn(letters, currentLetterIndex + 1, lower, math.Floor((lower + higher) / 2))
	} else {
		// Letter is R
		return findColumn(letters, currentLetterIndex + 1, math.Ceil((lower + higher)/2), higher)
	}
}

/**
 * findYourSeat takes a sorted list of all the seats and finds X where X is not found but X-1 and X+1 are
 */
func findYourSeatID(seats []int) int {
	for i := 1; i < len(seats) - 1; i++ {
		last := seats[i]
		next := seats[i+1]
		if next - last == 2 {
			return last + 1
		}
	}
	return -1
}
