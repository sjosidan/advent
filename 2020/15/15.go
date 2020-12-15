package main

import "fmt"

type Number struct {
	lastSpoken  []int
	timesSpoken int
}

func main() {
	//scanner := bufio.NewScanner(os.Stdin)

	//input := []int{0, 5, 4, 1, 10, 14, 7}
	gameMap := make(map[int]*Number)
	input := []int{1, 3, 2}
	n1 := Number{lastSpoken: []int{1}, timesSpoken: 1}
	gameMap[1] = &n1
	n2 := Number{lastSpoken: []int{2}, timesSpoken: 1}
	gameMap[3] = &n2
	n3 := Number{lastSpoken: []int{3}, timesSpoken: 1}
	gameMap[2] = &n3

	//input := []int{0, 5, 4, 1, 10, 14, 7}
	i := 4
	lastSpok := 2
	for {
		fmt.Println(i, lastSpok)

		if _, ok := gameMap[lastSpok]; !ok {

			if _, ok := gameMap[0]; ok {
				nx := Number{lastSpoken: []int{0}, timesSpoken: 1}

				gameMap[0] = &nx

			} else {
				gameMap[0].lastSpoken = append(gameMap[0].lastSpoken, i)
				gameMap[0].timesSpoken = len(gameMap[0].lastSpoken)

			}
			input = append(input, 0)

			lastSpok = 0
		} else {

			newVal := gameMap[lastSpok].lastSpoken[gameMap[lastSpok].timesSpoken-1] -
				gameMap[lastSpok].lastSpoken[gameMap[lastSpok].timesSpoken-1]

			if newV, ok := gameMap[lastSpok]; ok {
				newV.lastSpoken = append(gameMap[newVal].lastSpoken, i)
				newV.timesSpoken = len(gameMap[newVal].lastSpoken)
			} else {
				nx := Number{lastSpoken: []int{i}, timesSpoken: 1}
				gameMap[newVal] = &nx
			}
			input = append(input, newVal)
			lastSpok = newVal
		}
		i++
		if i == 2020 {
			break
		}
	}
	fmt.Println(input)
}
