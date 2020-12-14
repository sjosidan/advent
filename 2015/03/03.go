package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type SantaPosition struct {
	x int
	y int
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	santa := SantaPosition{}
	robo := SantaPosition{}
	santaHasVisited := make(map[string]bool)

	santaHasVisited["0-0"] = true
	santasTurn := true
	for scanner.Scan() {
		line := scanner.Text()
		for _, l := range line {
			switch string(l) {
			case "v":
				if santasTurn {
					santa.y--
					santasTurn = false
				} else {
					robo.y--
					santasTurn = true
				}
			case ">":
				if santasTurn {
					santa.x++
					santasTurn = false
				} else {
					robo.x++
					santasTurn = true
				}
			case "<":
				if santasTurn {
					santa.x--
					santasTurn = false
				} else {
					robo.x--
					santasTurn = true
				}
			case "^":
				if santasTurn {
					santa.y++
					santasTurn = false
				} else {
					robo.y++
					santasTurn = true
				}

			}
			houseKey := ""
			if santasTurn {
				houseKey = strconv.Itoa(santa.x) + "-" + strconv.Itoa(santa.y)
			} else {
				houseKey = strconv.Itoa(robo.x) + "-" + strconv.Itoa(robo.y)
			}

			santaHasVisited[houseKey] = true
		}
	}
	fmt.Println("Santa stops at", santa)
	fmt.Println("Santa stops at", len(santaHasVisited), " houses")
}
