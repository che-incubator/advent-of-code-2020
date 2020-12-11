package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	seats := readInput("./input")
	log.Println("How many seats end up occupied?")
	log.Println("Answer:", occupiedAfterStabilize(seats, step))
	log.Println("=============")
	log.Println("How many seats end up occupied?")
	log.Println("Answer:", occupiedAfterStabilize(seats, step2))
	log.Println("=============")
}

func occupiedAfterStabilize(seats [][]rune, stepFunc func([][]rune) ([][]rune, bool)) int {
	newSeats := seats
	for i := 0;; i = i + 1 {
		// printSeats(newSeats)
		var changed bool
		if newSeats, changed = stepFunc(newSeats); !changed {
			// printSeats(newSeats)
			return occupiedSeats(newSeats)
		}
	}
}

func occupiedSeats(seats [][]rune) int {
	occupied := 0
	for _, row := range seats {
		for _, seat := range row {
			if seat == '#' {
				occupied = occupied + 1
			}
		}
	}
	return occupied
}

func printSeats(seats [][]rune) {
	for _, row := range seats {		
		log.Println(string(row))
	}
	log.Println()
}

func step(seats [][]rune) (newSeats [][]rune, changed bool) {
	newSeats = make([][]rune, len(seats))

	for rowI, row := range seats {
		newSeats[rowI] = make([]rune, len(row))
		for colI, seat := range row {
			if seat == '.' {
				newSeats[rowI][colI] = seat
				continue
			} else {
				neighbors := occupiedNeghbors(seats, rowI, colI)
				if seat == 'L' && neighbors == 0 {
					newSeats[rowI][colI] = '#'
					changed = true
				} else if seat == '#' && neighbors >= 4 {
					newSeats[rowI][colI] = 'L'
					changed = true
				} else {
					newSeats[rowI][colI] = seat
				}
			}
		}
	}

	return newSeats, changed
}

func step2(seats [][]rune) (newSeats [][]rune, changed bool) {
	newSeats = make([][]rune, len(seats))

	for rowI, row := range seats {
		newSeats[rowI] = make([]rune, len(row))
		for colI, seat := range row {
			if seat == '.' {
				newSeats[rowI][colI] = seat
				continue
			} else {
				// log.Println("checking", rowI, colI)
				neighbors := occupiedNeghborsFar(seats, rowI, colI)
				// log.Println("neigh", neighbors)
				if seat == 'L' && neighbors == 0 {
					newSeats[rowI][colI] = '#'
					changed = true
				} else if seat == '#' && neighbors >= 5 {
					newSeats[rowI][colI] = 'L'
					changed = true
				} else {
					newSeats[rowI][colI] = seat
				}
			}
		}
	}

	return newSeats, changed
}

func occupiedNeghbors(seats [][]rune, row int, col int) int {
	occupiedNeigh := isOccupied(seats, row - 1, col - 1) + isOccupied(seats, row - 1, col) + isOccupied(seats, row - 1, col + 1)
 	occupiedNeigh = occupiedNeigh + isOccupied(seats, row, col - 1) + isOccupied(seats, row, col + 1)
	occupiedNeigh = occupiedNeigh + isOccupied(seats, row + 1, col - 1) + isOccupied(seats, row + 1, col) + isOccupied(seats, row + 1, col + 1)
	return occupiedNeigh
}

func occupiedNeghborsFar(seats [][]rune, row int, col int) int {
	occupiedNeigh := isOccupiedFar(seats, row, col, -1, -1) + isOccupiedFar(seats, row, col, -1, 0) + isOccupiedFar(seats, row, col, -1, 1)
 	occupiedNeigh = occupiedNeigh + isOccupiedFar(seats, row, col, 0, -1) + isOccupiedFar(seats, row, col, 0, 1)
	occupiedNeigh = occupiedNeigh + isOccupiedFar(seats, row, col, 1, -1) + isOccupiedFar(seats, row, col, 1, 0) + isOccupiedFar(seats, row, col, 1, 1)
	return occupiedNeigh
}

func isOccupiedFar(seats [][]rune, row int, col int, dRow int, dCol int) int {
	for rowI, colI := row + dRow, col + dCol; ; rowI, colI = rowI + dRow, colI + dCol {
			// log.Println("occupied? ", rowI, colI)
			if outOfBounds(seats, rowI, colI) {
				// log.Println("out of bounds. next. ", rowI, colI)
				return 0
			} else if seats[rowI][colI] == 'L' || seats[rowI][colI] == '#' {
				return isOccupied(seats, rowI, colI)
			}
	} 
}

func outOfBounds(seats [][]rune, row int, col int) bool {
	if row < 0 || col < 0 {
		return true
	}
	if row >= len(seats) || col >= len(seats[0]) {
		return true
	}
	return false
}

func isOccupied(seats [][]rune, row int, col int) int {
	if outOfBounds(seats, row, col) {
		return 0
	}
	if seats[row][col] == '#' {
		return 1
	}
	return 0
}

func readInput(input string) [][]rune {
	seats := make([][]rune, 0)

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		row := make([]rune, 0)
		for _, c := range line {
			row = append(row, c)
		}
		seats = append(seats, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return seats
}
