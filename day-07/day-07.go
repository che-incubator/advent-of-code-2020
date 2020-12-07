package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	bags := readBags("./input")	// map[bag] is in map[bag] amount times
	// log.Println(len(bags))
	validColors := validColors(bags, "shiny gold", make(map[string]bool)) - 1	// -1 because we dont count box passed here
	log.Println(validColors)

	bags2 := readBags2("./input")
	// log.Println(bags2)
	totalBags := totalBags(bags2, "shiny gold") - 1 // -1 dont count itself
	log.Println(totalBags)
}

func totalBags(bags map[string]map[string]int, color string) int {
	count := 0
	if len(bags[color]) == 0 {
		// log.Println("no more bags", color)
		return 1
	} else {
		for b, a := range bags[color] {
			// log.Printf("%s contains %d of %s\n", color, a, b)
			count = count + (a * totalBags(bags, b))
		}
	}
	return count + 1
}

func validColors(bags map[string]map[string]int, color string, counted map[string]bool) int {
	count := 0
	if _, ok := bags[color]; ok {
		for k, _ := range bags[color] {
			count = count + validColors(bags, k, counted)
			if _, alreadyCounted := counted[color]; !alreadyCounted {
				counted[color] = true
				count = count + 1
			}
		}
	} else {
		if _, alreadyCounted := counted[color]; alreadyCounted {
			return 0
		} else {
			counted[color] = true
			return 1
		}
	}
	return count
}

func readBags(input string) map[string]map[string]int {
	bags := make(map[string]map[string]int)

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		bag, contains := parse(scanner.Text())
		// log.Println(bag, contains)

		// log.Println(bag)
		for cBag, cAmount := range contains {
			// log.Println(cBag, cAmount)
			if _, ok := bags[cBag]; ok {
				bags[cBag][bag] = cAmount
			} else {
				bags[cBag] = make(map[string]int)
				bags[cBag][bag] = cAmount
			}
		}
	}
	
	return bags
}

func readBags2(input string) map[string]map[string]int {
	bags := make(map[string]map[string]int)

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		bag, contains := parse(scanner.Text())

		for cBag, cAmount := range contains {
			if _, ok := bags[bag]; !ok {
				bags[bag] = make(map[string]int)
			}
			bags[bag][cBag] = cAmount
		}
	}
	
	return bags
}

func parse(line string) (string, map[string]int) {
	splitted := strings.Split(line, "bags contain")
	contains := strings.Split(splitted[1], "bag")

	bags := make(map[string]int)
	
	for _, containsBags := range contains {
		containsBags = strings.TrimLeft(containsBags, "s")
		containsBags = strings.Trim(containsBags, " ,.")
		bagsSplitted := strings.SplitN(containsBags, " ", 2)
		if len(bagsSplitted) == 2 && bagsSplitted[0] != "no" {
			amount, err := strconv.Atoi(bagsSplitted[0])
			if err != nil {
				log.Println(bagsSplitted[0])
				log.Println(bagsSplitted[1])
				log.Fatal(err)
			}
			bags[bagsSplitted[1]] = amount
		}
	}
	return strings.TrimSpace(splitted[0]), bags
}
