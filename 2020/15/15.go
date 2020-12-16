package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	lastSpoken  []int
	timesSpoken int
}

var input []int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	gameMap := make(map[int]*Number)

	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		ax := strings.Split(line, ",")
		for _, a := range ax {
			z, err := strconv.Atoi(a)
			if err == nil {
				if _, ok := gameMap[z]; ok {
					input = append(input, z)
					gameMap[z].lastSpoken = append(gameMap[z].lastSpoken, i)
					gameMap[z].timesSpoken = len(gameMap[z].lastSpoken)
					i = i + 1
				} else {
					input = append(input, z)
					nx := Number{lastSpoken: []int{i}, timesSpoken: 1}
					gameMap[z] = &nx
					i = i + 1
				}
			}

		}
	}
	fmt.Println(input)
	for k, v := range gameMap {
		fmt.Println(k, v.lastSpoken)
	}
	fmt.Println("")
	fmt.Println("----")
	i = len(input)
	i = i + 1
	lastSpok := input[len(input)-1]
	for {

		if len(gameMap[lastSpok].lastSpoken) == 1 {

			if _, ok := gameMap[0]; !ok {
				nx := Number{lastSpoken: []int{i}, timesSpoken: 1}

				gameMap[0] = &nx

			} else {
				gameMap[0].lastSpoken = append(gameMap[0].lastSpoken, i)
				gameMap[0].timesSpoken = len(gameMap[0].lastSpoken)

			}
			input = append(input, 0)

			lastSpok = 0
		} else if len(gameMap[lastSpok].lastSpoken) > 1 {
			newVal := gameMap[lastSpok].lastSpoken[len(gameMap[lastSpok].lastSpoken)-1] -
				gameMap[lastSpok].lastSpoken[len(gameMap[lastSpok].lastSpoken)-2]

			if _, ok := gameMap[newVal]; ok {

				gameMap[newVal].lastSpoken = append(gameMap[newVal].lastSpoken, i)
				gameMap[newVal].timesSpoken = len(gameMap[newVal].lastSpoken)

			} else {
				nx := Number{lastSpoken: []int{i}, timesSpoken: 1}
				gameMap[newVal] = &nx
			}
			input = append(input, newVal)
			lastSpok = newVal
		} else {
			fmt.Println("shouldnt be here")
		}

		i = i + 1
		if i > 30000002 {
			break
		}

	}
	fmt.Println("A", input[2020-1])

	fmt.Println("B", input[30000000-1])

}
