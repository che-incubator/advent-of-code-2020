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
	processShipLocation("./input")
	processWaypointLocation("./input")
}

func processShipLocation(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	currDirection := "E"
	x := 0
	y := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		dir := string(line[0])
		mov, _ := strconv.Atoi(line[1:])

		switch dir {
		case "N":
			{
				y += mov
			}
		case "S":
			{
				y -= mov
			}
		case "E":
			{
				x += mov
			}
		case "W":
			{
				x -= mov
			}
		case "L":
			{
				currDirection = generateRotation(false, currDirection, mov)
			}
		case "R":
			{
				currDirection = generateRotation(true, currDirection, mov)
			}
		case "F":
			{
				if currDirection == "N" {
					y += mov
				} else if currDirection == "S" {
					y -= mov
				} else if currDirection == "E" {
					x += mov
				} else if currDirection == "W" {
					x -= mov
				}
			}
		}

	}

	m := Abs(x) + Abs(y)
	fmt.Printf("P1: %d\n", m)
}

func processWaypointLocation(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	x := 0
	y := 0
	waypointX := 10
	waypointY := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		dir := string(line[0])
		mov, _ := strconv.Atoi(line[1:])

		switch dir {
		case "N":
			{
				waypointY += mov
			}
		case "S":
			{
				waypointY -= mov
			}
		case "E":
			{
				waypointX += mov
			}
		case "W":
			{
				waypointX -= mov
			}
		case "L":
			{
				waypointX, waypointY = generateWaypointRotation(false, mov, waypointX, waypointY)
			}
		case "R":
			{
				waypointX, waypointY = generateWaypointRotation(true, mov, waypointX, waypointY)
			}
		case "F":
			{
				x += mov * waypointX
				y += mov * waypointY
			}
		}

	}

	m := Abs(x) + Abs(y)
	fmt.Printf("P2: %d\n", m)
}

func generateWaypointRotation(isRight bool, rotation, x, y int) (int, int) {
	if isRight {
		if rotation == 90 {
			return y, -x
		} else if rotation == 180 {
			return -x, -y
		} else if rotation == 270 {
			return -y, x
		} else {
			return x, y
		}
	} else {
		if rotation == 90 {
			return -y, x
		} else if rotation == 180 {
			return -x, -y
		} else if rotation == 270 {
			return y, -x
		} else {
			return x, y
		}
	}
}

func generateRotation(isRight bool, currDirection string, rotation int) string {
	possible := "ESWN"
	possibleLen := len(possible)
	if isRight {
		if rotation == 90 {
			return string(possible[(strings.Index(possible, currDirection)+1)%possibleLen])
		} else if rotation == 180 {
			return string(possible[(strings.Index(possible, currDirection)+2)%possibleLen])
		} else if rotation == 270 {
			return string(possible[(strings.Index(possible, currDirection)+3)%possibleLen])
		} else {
			return currDirection
		}
	} else {
		if rotation == 90 {
			return string(possible[(strings.Index(possible, currDirection)+possibleLen-1)%possibleLen])
		} else if rotation == 180 {
			return string(possible[(strings.Index(possible, currDirection)+possibleLen-2)%possibleLen])
		} else if rotation == 270 {
			return string(possible[(strings.Index(possible, currDirection)+possibleLen-3)%possibleLen])
		} else {
			return currDirection
		}
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
