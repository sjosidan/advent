package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	insA := strings.Split(line, ",")

	scanner.Scan()
	line = scanner.Text()
	insB := strings.Split(line, ",")

	thisMap := make(map[string]int)

	parse(thisMap, insA)
	parseb(thisMap, insB)
}

func parse(a map[string]int, b []string) {
	currX := 0
	currY := 0
	steps := 0
	for _, instr := range b {
		length, _ := strconv.Atoi(instr[1:])
		instruction := instr[0:1]

		switch instruction {
		case "L":
			for i := 0; i < length; i++ {
				currX = currX - 1
				z := strconv.Itoa(currX) + "-" + strconv.Itoa(currY)
				steps = steps + 1

				a[z] = steps
			}
		case "R":
			for i := 0; i < length; i++ {
				currX = currX + 1
				z := strconv.Itoa(currX) + "-" + strconv.Itoa(currY)
				steps = steps + 1

				a[z] = steps

			}
		case "U":
			for i := 0; i < length; i++ {
				currY = currY + 1
				z := strconv.Itoa(currX) + "-" + strconv.Itoa(currY)
				steps = steps + 1

				a[z] = steps

			}
		case "D":
			for i := 0; i < length; i++ {
				currY = currY - 1
				z := strconv.Itoa(currX) + "-" + strconv.Itoa(currY)
				steps = steps + 1

				a[z] = steps

			}
		}

	}
}

func parseb(a map[string]int, b []string) {
	maxHit := float64(100000000)
	currX := 0
	currY := 0
	steps := 0
	minSteps := 10000000
	for _, instr := range b {
		length, _ := strconv.Atoi(instr[1:])
		instruction := instr[0:1]

		switch instruction {
		case "L":
			for i := 0; i < length; i++ {
				currX = currX - 1
				z := strconv.Itoa(currX) + "-" + strconv.Itoa(currY)
				steps = steps + 1

				if _, ok := a[z]; ok {
					if maxHit > abss(currX, currY) {
						maxHit = abss(currX, currY)
					}
					if minSteps > steps+a[z] {
						minSteps = steps + a[z]
					}
				}
			}
		case "R":
			for i := 0; i < length; i++ {
				currX = currX + 1
				z := strconv.Itoa(currX) + "-" + strconv.Itoa(currY)
				steps = steps + 1
				if _, ok := a[z]; ok {

					if maxHit > abss(currX, currY) {
						maxHit = abss(currX, currY)

					}
					if minSteps > steps+a[z] {
						minSteps = steps + a[z]
					}
				}
			}
		case "U":
			for i := 0; i < length; i++ {
				currY = currY + 1
				z := strconv.Itoa(currX) + "-" + strconv.Itoa(currY)
				steps = steps + 1
				if _, ok := a[z]; ok {

					if maxHit > abss(currX, currY) {
						maxHit = abss(currX, currY)

					}
					if minSteps > steps+a[z] {
						minSteps = steps + a[z]
					}
				}
			}
		case "D":
			for i := 0; i < length; i++ {
				currY = currY - 1
				z := strconv.Itoa(currX) + "-" + strconv.Itoa(currY)
				steps = steps + 1
				if _, ok := a[z]; ok {
					if maxHit > abss(currX, currY) {
						maxHit = abss(currX, currY)

					}
					if minSteps > steps+a[z] {
						minSteps = steps + a[z]
					}
				}
			}
		}

	}
	fmt.Println(maxHit)
	fmt.Println(minSteps)

}

func abss(a int, b int) (result float64) {

	result = math.Abs(float64(a)) + math.Abs(float64(b))
	return
}
