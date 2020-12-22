package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var player1deck []int
	var player2deck []int
	playerone := false
	first := true
	for scanner.Scan() {
		line := scanner.Text()
		val, notNum := strconv.Atoi(line)
		if notNum != nil {

			if line == "Player 1:" {
				playerone = true
			}
			if line == "Player 2:" {
				playerone = false
			}
		} else {

			if playerone && !first {
				player1deck = append(player1deck, val)

			}
			if !playerone && !first {
				val, _ := strconv.Atoi(line)
				player2deck = append(player2deck, val)

			}

		}
		first = false
	}

	fmt.Println("Player 1 deck", player1deck)
	fmt.Println("Player 2 deck", player2deck)

	p1winner := false
	haveWinner := false
	for !haveWinner {

		p1, p2 := play(player1deck, player2deck)
		if len(p1) == 0 {
			fmt.Println("P2 wins")
			p1winner = false
			haveWinner = true

		}
		if len(p2) == 0 {
			fmt.Println("P1 wins")
			p1winner = true
			haveWinner = true

		}
		player1deck = p1
		player2deck = p2

	}
	fmt.Println("== Post-game results ==")
	fmt.Println("Player 1's deck: ", player1deck)
	fmt.Println("Player 2's deck: ", player2deck)
	totalScore := 0
	if p1winner {
		for i := len(player1deck) - 1; i > -1; i-- {
			multiplier := (len(player1deck) - i)
			totalScore = totalScore + (player1deck[i] * multiplier)
		}
	} else {
		for i := len(player2deck) - 1; i > -1; i-- {
			multiplier := (len(player2deck) - i)
			fmt.Println(player2deck[i], multiplier)
			totalScore = totalScore + (player2deck[i] * multiplier)
		}
	}
	fmt.Println("total score", totalScore)
}

func play(p1 []int, p2 []int) (res1 []int, res2 []int) {
	res1 = p1[1:]
	res2 = p2[1:]

	if p1[0] > p2[0] {
		fmt.Println("Player 1 wins round!")
		res1 = append(res1, p1[0], p2[0])
	} else {
		fmt.Println("Player 2 wins round!")
		res2 = append(res2, p2[0], p1[0])
	}
	return
}
