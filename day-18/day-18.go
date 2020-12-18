package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var ch string
var pos int
var line string
var isP2 bool

func main() {
	fileLocation := "./input"
	part1(fileLocation)
	part2(fileLocation)
}

func part1(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalValue := 0
	for scanner.Scan() {
		line = scanner.Text()
		totalValue += parse()
	}
	fmt.Printf("Using the initial rules: %d\n", totalValue)
}

func part2(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	isP2 = true
	scanner := bufio.NewScanner(file)
	totalValue := 0
	for scanner.Scan() {
		line = scanner.Text()
		totalValue += parse()
	}
	fmt.Printf("Using the new rules: %d\n", totalValue)
}

func nextChar() {
	pos += 1
	if pos < len(line) {
		ch = string(line[pos])
	}
}

func eatNext(toEat string) bool {
	for ch == " " {
		nextChar()
	}
	if ch == toEat {
		nextChar()
		return true
	}
	return false
}

func parse() int {
	pos = -1
	nextChar()
	var value int
	for pos < len(line) {
		ch = string(line[pos])
		value = parseExpression()
	}
	return value
}

func parseExpression() int {
	var x int
	if isP2 {
		x = parseTerm()
		for {
			if eatNext("*") {
				x *= parseTerm()
			} else if eatNext("+") {
				x += parseTerm()
			} else {
				return x
			}
		}
	} else {
		x = parseFactor()
		for {
			if eatNext("*") {
				x *= parseFactor()
			} else if eatNext("+") {
				x += parseFactor()
			} else {
				return x
			}
		}
	}
}

func parseTerm() int {
	x := parseFactor()
	for {
		if eatNext("+") {
			x += parseFactor()
		} else {
			return x
		}
	}
}

func isEnd() bool {
	if pos >= len(line) {
		return true
	}
	return false
}

func parseFactor() int {
	if eatNext("+") {
		return parseFactor()
	}

	var x int
	startPos := pos
	if eatNext("(") {
		x = parseExpression()
		eatNext(")")
	} else {
		// We should have a number, parse until we hit a space
		for (ch >= "0" && ch <= "9") && !isEnd() {
			nextChar()
		}
		x, _ = strconv.Atoi(line[startPos:pos])
	}

	return x
}
