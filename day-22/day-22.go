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
	player1Cards, player2Cards := parseInput()
	combat(player1Cards, player2Cards)
	state := recursiveCombat(player1Cards, player2Cards)
	total := sumCards(state.winningCards)
	fmt.Printf("The total score for part 2 is: %d\n", total)
}

func parseInput() ([]int, []int) {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	spacesFound := false
	player1Cards := make([]int, 0)
	player2Cards := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			spacesFound = true
		}

		if spacesFound == false && line != "Player 1:" {
			c, _ := strconv.Atoi(line)
			player1Cards = append(player1Cards, c)
		} else if spacesFound == true && line != "Player 2:" && line != "" {
			c, _ := strconv.Atoi(line)
			player2Cards = append(player2Cards, c)
		}

	}
	return player1Cards, player2Cards
}

func combat(player1Cards []int, player2Cards []int) {
	player1Index := 0
	player2Index := 0
	for len(player1Cards) != 0 && len(player2Cards) != 0 {

		if player1Cards[player1Index] > player2Cards[player2Index] {
			firstCard := player1Cards[player1Index]
			player1Cards = removeIndex(player1Cards, 0)
			player1Cards = append(player1Cards, firstCard)
			player1Cards = append(player1Cards, player2Cards[player2Index])
			player2Cards = removeIndex(player2Cards, 0)
		} else if player1Cards[player1Index] < player2Cards[player2Index] {
			firstCard := player2Cards[player2Index]
			player2Cards = removeIndex(player2Cards, 0)
			player2Cards = append(player2Cards, firstCard)
			player2Cards = append(player2Cards, player1Cards[player1Index])
			player1Cards = removeIndex(player1Cards, 0)
		}

	}

	var totalScore int
	if len(player1Cards) == 0 {
		// player 2 won
		totalScore = sumCards(player2Cards)
	} else {
		// player 1 won
		totalScore = sumCards(player1Cards)
	}
	fmt.Printf("The total score for part 1 is: %d\n", totalScore)
}

type State struct {
	winner       int
	winningCards []int
}

func recursiveCombat(player1Cards []int, player2Cards []int) State {
	player1Index := 0
	player2Index := 0
	playedStates := make(map[string]int)
	for len(player1Cards) != 0 && len(player2Cards) != 0 {

		currentState := fmt.Sprintf("player1:%s player2: %s", intToString(player1Cards), intToString(player2Cards))
		if _, ok := playedStates[currentState]; ok {
			return State{
				winner:       1,
				winningCards: player1Cards,
			}
		}
		playedStates[currentState] = 0

		firstPlayer1Card := player1Cards[player1Index]
		firstPlayer2Card := player2Cards[player2Index]
		player1Cards = removeIndex(player1Cards, 0)
		player2Cards = removeIndex(player2Cards, 0)

		var winner int
		if len(player1Cards) >= firstPlayer1Card && len(player2Cards) >= firstPlayer2Card {
			state := recursiveCombat(player1Cards[0:firstPlayer1Card], player2Cards[0:firstPlayer2Card])
			winner = state.winner
		} else if firstPlayer1Card > firstPlayer2Card {
			winner = 1
		} else {
			winner = 2
		}

		if winner == 1 {
			player1Cards = append(player1Cards, firstPlayer1Card)
			player1Cards = append(player1Cards, firstPlayer2Card)
		} else {
			player2Cards = append(player2Cards, firstPlayer2Card)
			player2Cards = append(player2Cards, firstPlayer1Card)
		}

	}

	var winner int
	var winningCards []int
	if len(player1Cards) > 0 {
		winner = 1
		winningCards = player1Cards
	} else {
		winner = 2
		winningCards = player2Cards
	}

	return State{
		winner:       winner,
		winningCards: winningCards,
	}
}

func sumCards(cards []int) int {
	total := 0
	soFar := 1
	for i := len(cards) - 1; i >= 0; i-- {
		total += (cards[i] * soFar)
		soFar += 1
	}
	return total
}

func intToString(a []int) string {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}

	return strings.Join(b, ",")
}

func removeIndex(s []int, index int) []int {
	cop := copy(s)
	if index == len(s) {
		return s
	}
	return append(cop[:index], cop[index+1:]...)
}

func copy(myArr []int) []int {
	newArr := make([]int, 0)
	for _, v := range myArr {
		newArr = append(newArr, v)
	}
	return newArr
}
