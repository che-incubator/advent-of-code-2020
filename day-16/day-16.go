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
	readInput("./input")
}

func readInput(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tickets := make(map[string]func(int) bool)
	scanner := bufio.NewScanner(file)
	numSpaces := 0
	invalidTickets := make([]int, 0)
	validTickets := make([][]int, 0)
	yourTicket := make([]int, 0)
	arrayOfMapToPossibleSolutions := make([]map[int][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			numSpaces += 1
		}
		if numSpaces == 0 {
			fieldItem := strings.Split(line, ":")
			ticketName := fieldItem[0]
			boundariesSplit := strings.Split(strings.Trim(fieldItem[1], " "), " or ")
			leftBoundary := strings.Split(boundariesSplit[0], "-")
			rightBoundary := strings.Split(boundariesSplit[1], "-")

			tickets[ticketName] = func(toEvaluate int) bool {
				firstLeft, _ := strconv.Atoi(leftBoundary[0])
				firstRight, _ := strconv.Atoi(leftBoundary[1])
				secondLeft, _ := strconv.Atoi(rightBoundary[0])
				secondRight, _ := strconv.Atoi(rightBoundary[1])
				return (toEvaluate >= firstLeft && toEvaluate <= firstRight) || (toEvaluate >= secondLeft && toEvaluate <= secondRight)
			}
		} else if numSpaces == 1 && line != "your ticket:" && line != "" {
			yourTicket = strToIntArray(line)
		} else if numSpaces == 2 && line != "nearby tickets:" && line != "" {
			arr := strToIntArray(line)
			potentialInvalidTickets, possibleSolutions := checkIfValid(arr, tickets)
			if len(potentialInvalidTickets) == 0 {
				validTickets = append(validTickets, arr)
				arrayOfMapToPossibleSolutions = append(arrayOfMapToPossibleSolutions, possibleSolutions)
			} else {
				invalidTickets = append(invalidTickets, potentialInvalidTickets...)
			}
		}
	}
	sumInvalid := 0
	for _, val := range invalidTickets {
		sumInvalid += val
	}
	fmt.Printf("There sum of invalid tickets is: %d\n", sumInvalid)

	total := intersection(arrayOfMapToPossibleSolutions, yourTicket)
	fmt.Printf("The departure totals multiplied are: %d\n", total)
}

// Check if a given line is valid according to a map of evaluators
func checkIfValid(line []int, evaluators map[string]func(int) bool) ([]int, map[int][]string) {
	valid := false
	invalidNums := make([]int, 0)
	validSolutionForItem := make(map[int][]string)
	for i, intToCheck := range line {
		for evalName, evalFunc := range evaluators {
			if evalFunc(intToCheck) {
				valid = true
				validSolutionForItem[i] = append(validSolutionForItem[i], evalName)
			}
		}
		if !valid {
			invalidNums = append(invalidNums, intToCheck)
		}
		valid = false
	}
	return invalidNums, validSolutionForItem
}

// Get the intersection of all the columns and map them to a rule
func intersection(possibleSolutions []map[int][]string, yourTicket []int) int {
	intersectionMap := buildIntersectionMap(possibleSolutions)
	rules := getRules(intersectionMap, len(possibleSolutions))
	matches := eliminateRules(rules)
	total := multiplyDepartures(matches, yourTicket)
	return total
}

// Build a mapping of all lines to their columns to their possible solutions
func buildIntersectionMap(possibleSolutions []map[int][]string) map[int]map[string]int {
	intersectionMap := make(map[int]map[string]int)
	for _, solutionMaps := range possibleSolutions {
		for index, solutionForIndex := range solutionMaps {
			for _, field := range solutionForIndex {
				_, ok := intersectionMap[index]
				if ok {
					intersectionMap[index][field] = intersectionMap[index][field] + 1
				} else {
					intersectionMap[index] = make(map[string]int)
					_, ok2 := intersectionMap[index][field]
					if ok2 {
						intersectionMap[index][field] = 1
					} else {
						intersectionMap[index][field] = 1
					}
				}
			}

		}
	}
	return intersectionMap
}

// Get all the possible rules by eliminating possible solutions that don't match every neighbouring ticket index for a
// given index
func getRules(intersectionMap map[int]map[string]int, possibleSolutionsCount int) map[int][]string {
	rules := make(map[int][]string)
	for i, v := range intersectionMap {
		for name, numOccurences := range v {
			// If the number of occurrences is equal to the total number of lines then we can consider it a rule
			if numOccurences == possibleSolutionsCount {
				rules[i] = append(rules[i], name)
			}
		}
	}
	return rules
}

// Keep eliminating every rule that has 1 item until there are none left
// Return a mapping of index to the ticket field
func eliminateRules(rules map[int][]string) map[int]string {
	matches := make(map[int]string)
	for len(matches) < len(rules) {
		var match string
		for int, value := range rules {
			if len(value) == 1 {
				matches[int] = value[0]
				match = value[0]
				rules[int] = remove(value, match)
				break
			}
		}

		for int, value := range rules {
			rules[int] = remove(value, match)
		}
	}
	return matches
}

// Sum all the departures on your ticket
func multiplyDepartures(matches map[int]string, yourTicket []int) int {
	total := 1
	for i, v := range matches {
		if strings.HasPrefix(v, "departure") {
			total = total * yourTicket[i]
		}
	}
	return total
}

// Copy myMap into a new map of the same type
func copy(myMap map[string]func(int) bool) map[string]func(int) bool {
	newMap := make(map[string]func(int) bool)
	for k, v := range myMap {
		newMap[k] = v
	}
	return newMap
}

// Remove toDelete from s if found
func remove(s []string, toDelete string) []string {
	for index, val := range s {
		if val == toDelete {
			return append(s[:index], s[index+1:]...)
		}
	}
	return s
}

// Convert a line to an int array
func strToIntArray(line string) []int {
	parts := strings.Split(line, ",")
	yourTicket := make([]int, len(parts), len(parts))
	for i := 0; i < len(parts); i++ {
		yourTicket[i], _ = strconv.Atoi(parts[i])
	}
	return yourTicket
}
