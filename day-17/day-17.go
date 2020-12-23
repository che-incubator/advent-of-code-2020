package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	template = "w:%d:x:%d:y:%d:z:%d"
)

func parseInput() (int, int, map[string]bool) {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	activePositions := map[string]bool{}
	var inputHeight int
	var inputLength int
	for scanner.Scan() {
		line := scanner.Text()
		inputLength = len(line)
		inputHeight += 1

		for loc, c := range line {
			if string(c) == "#" {
				activePositions[fmt.Sprintf(template, 0, loc, inputHeight, 0)] = true
			}
		}
	}
	return inputHeight, inputLength, activePositions
}

func copy(myMap map[string]bool) map[string]bool {
	newMap := make(map[string]bool)
	for k, v := range myMap {
		newMap[k] = v
	}
	return newMap
}

func main() {
	inputHeight, inputLength, activePositions := parseInput()
	startingStateP1 := copy(activePositions)
	startingStateP2 := copy(activePositions)

	for j := 1; j <= 6; j++ {
		nextStateP1 := map[string]bool{}
		nextStateP2 := map[string]bool{}

		// Part 1
		for x := 0 - j; x <= 0+j; x++ {
			for y := 0 - j; y <= inputLength+j; y++ {
				for z := 0 - j; z <= inputHeight+j; z++ {
					newState := nextState(0, y, z, x, startingStateP1)
					nextStateP1 = setNextStatePosition(0, y, z, x, nextStateP1, newState)
				}
			}
		}

		// Part 2
		for w := 0 - j; w <= 0+j; w++ {
			for x := 0 - j; x <= 0+j; x++ {
				for y := 0 - j; y <= inputHeight+j; y++ {
					for z := 0 - j; z <= inputHeight+j; z++ {
						newState := nextState(w, y, z, x, startingStateP2)
						nextStateP2 = setNextStatePosition(w, y, z, x, nextStateP2, newState)
					}
				}
			}
		}

		startingStateP1 = nextStateP1
		startingStateP2 = nextStateP2
	}
	fmt.Printf("Part 1: %d\n", len(startingStateP1))
	fmt.Printf("Part 2: %d\n", len(startingStateP2))
}

func findTotalActiveNumbers(w, x, y, z int, state map[string]bool) int {
	totalActiveNumbers := 0
	for a := x - 1; a <= x+1; a++ {
		for b := y - 1; b <= y+1; b++ {
			for c := z - 1; c <= z+1; c++ {
				for d := w - 1; d <= w+1; d++ {
					if d == w && a == x && b == y && c == z {
						continue
					}
					if isCurrentPosActive(d, a, b, c, state) {
						totalActiveNumbers++
					}
				}
			}
		}
	}
	return totalActiveNumbers
}

func nextState(w, x, y, z int, state map[string]bool) bool {
	totalActiveNumbers := findTotalActiveNumbers(w, x, y, z, state)
	currentPosIsActive := isCurrentPosActive(w, x, y, z, state)

	// Check if the current position has 2 or 3 neighbours that are active
	if currentPosIsActive && (totalActiveNumbers == 2 || totalActiveNumbers == 3) {
		return true
	} else if !currentPosIsActive && totalActiveNumbers == 3 {
		return true
	} else {
		return false
	}
}

func isCurrentPosActive(w, x, y, z int, state map[string]bool) bool {
	if _, ok := state[fmt.Sprintf(template, w, x, y, z)]; ok {
		return state[fmt.Sprintf(template, w, x, y, z)]
	} else {
		return false
	}
}

func setNextStatePosition(w, x, y, z int, state map[string]bool, newVal bool) map[string]bool {
	if newVal {
		state[fmt.Sprintf(template, w, x, y, z)] = true
	} else {
		delete(state, fmt.Sprintf(template, w, x, y, z))
	}
	return state
}
