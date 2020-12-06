package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	readCustomsCard("./input")
}

func readCustomsCard(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	newCustomsCard := make(map[string]int) // An individual customs card
	totalYes := 0 // The total number of yes per custom card
	everyoneSaidYes := 0 // The questions to which everyone on the same card answered "yes" to
	numPeopleOnCard := 0 // The total number of people on one declaration card

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			totalYes += len(newCustomsCard)
			everyoneSaidYes += findEveryoneSaidYes(newCustomsCard, numPeopleOnCard)

			// Re-set the map and start verifying a new customs card
			newCustomsCard = make(map[string]int)
			numPeopleOnCard = 0
		} else {
			// Iterate over the card declaration and keep track of how many times a question was answered
			for _, question := range line {
				q := string(question)
				val, ok := newCustomsCard[q]
				if ok {
					newCustomsCard[q] = val + 1
				} else {
					newCustomsCard[q] = 1
				}
			}
			numPeopleOnCard += 1
		}
	}

	totalYes += len(newCustomsCard)
	everyoneSaidYes += findEveryoneSaidYes(newCustomsCard, numPeopleOnCard)

	fmt.Printf("There are %d total yes\n", totalYes)
	fmt.Printf("There are %d total questions where everyone on one card said yes\n", everyoneSaidYes)
}

/**
 * Find questions to which everyone on the same card answered "yes" to
 */
func findEveryoneSaidYes(m map[string]int, numPeople int) int {
	count := 0
	for _, v := range m {
		if v == numPeople {
			count += 1
		}
	}
	return count
}
