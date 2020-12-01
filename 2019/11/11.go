package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var firstOneRunning = false
var globalMax = 0

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	a := strings.Split(line, ",")
	orig := make([]int, len(a))
	for i, v := range a {
		orig[i], _ = strconv.Atoi(v)

	}

	//arr := []int{0, 1, 2, 3, 4}
	//arr := []int{5, 6, 7, 8, 9}
	//for _, a := range permutations(arr) {
	wg := new(sync.WaitGroup)
	b := append(orig[:0:0], orig...)
	memory := make([]int, 200000)
	for i := 0; i < 200000; i++ {
		memory[i] = 0
	}
	b = append(b, memory...)
	wg.Add(1)
	input1 := make(chan int)

	output := make(chan int)
	//fmt.Println(b)
	run(b, &input1, &output, "COMP1", wg)
	//input1 <- 1
	//fmt.Println("out", <-output)
	wg.Wait()

	close(input1)
	close(output)

	//}

}

func run(b []int, inputChannel *chan int, outputChannel *chan int, name string, wg *sync.WaitGroup) int {
	painted := 0
	shipHull := [150][150]int{}
	for i := 0; i < 150; i++ {
		for j := 0; j < 150; j++ {
			shipHull[i][j] = 0
		}
	}
	shipHullPaintCount := [150][150]int{}
	for i := 0; i < 150; i++ {
		for j := 0; j < 150; j++ {
			shipHullPaintCount[i][j] = 0
		}
	}
	positionX := 75
	positionY := 75
	shipHull[positionX][positionY] = 1
	paint := true
	direction := "N"
	defer wg.Done()
	if name == "COMP1" {
		firstOneRunning = true
	}
	outputParam := 0
	i := 0
	relativeBase := 0
	for i < len(b) {

		op := b[i] % 100
		modeA := b[i] / 10000
		modeB := b[i] / 1000 % 10
		modeC := b[i] / 100 % 10
		switch op {
		case 1:
			//add
			//fmt.Println(b[i], b[i+1], b[i+2], b[i+3])
			first := getParamterOutput(modeC, relativeBase, i, b, 1)
			second := getParamterOutput(modeB, relativeBase, i, b, 2)
			third := getParamterOutput(modeA, relativeBase, i, b, 3)
			b[third] = b[first] + b[second]
			i = i + 4

		case 2:
			//Multiply
			//fmt.Println(b[i], b[i+1], b[i+2], b[i+3])
			first := getParamterOutput(modeC, relativeBase, i, b, 1)
			second := getParamterOutput(modeB, relativeBase, i, b, 2)
			third := getParamterOutput(modeA, relativeBase, i, b, 3)

			b[third] = b[first] * b[second]

			i = i + 4
		case 3:
			//fmt.Println(b[i], b[i+1])
			//inputStuff := <-*inputChannel

			//first := getParamterOutput(modeC, relativeBase, i, b, 1)
			b[b[i+1]+relativeBase] = shipHull[positionX][positionY]
			i = i + 2

		case 4:
			//fmt.Println(b[i], b[i+1])
			color := "white"
			first := getParamterOutput(modeC, relativeBase, i, b, 1)
			outputParam = b[first]

			//b[first] = b[i+1]
			if paint {
				if b[first] == 0 {
					color = "black"
					shipHull[positionX][positionY] = 0
				} else {
					shipHull[positionX][positionY] = 1
				}
				paint = false
				shipHullPaintCount[positionX][positionY] = shipHullPaintCount[positionX][positionY] + 1
				if shipHullPaintCount[positionX][positionY] == 1 {
					painted = painted + 1
				}
				fmt.Println("Paint", color, "on", positionX, positionY)

			} else {
				paint = true
				switch direction {
				case "S":
					switch b[first] {
					case 0:
						direction = "E"
						positionX = positionX + 1
					case 1:
						direction = "W"
						positionX = positionX - 1
					}
				case "N":
					switch b[first] {
					case 0:
						direction = "W"
						positionX = positionX - 1
					case 1:
						direction = "E"
						positionX = positionX + 1
					}
				case "W":
					switch b[first] {
					case 0:
						direction = "S"
						positionY = positionY - 1
					case 1:
						direction = "N"
						positionY = positionY + 1
					}
				case "E":
					switch b[first] {
					case 0:
						direction = "N"
						positionY = positionY + 1
					case 1:
						direction = "S"
						positionY = positionY - 1
					}
				}
			}

			if name == "COMP5" {
				if !firstOneRunning {
					fmt.Println(name, "Done with last one", b[b[i+1]])
					if globalMax < b[first] {
						globalMax = b[first]
					}
				} else {
					//	*outputChannel <- b[first]
				}

			}
			i = i + 2

		case 5:
			//fmt.Println(b[i], b[i+1], b[i+2])

			//JUMP if true
			first := getParamterOutput(modeC, relativeBase, i, b, 1)

			second := getParamterOutput(modeB, relativeBase, i, b, 2)
			if b[first] != 0 {
				i = b[second]
			} else {
				i = i + 3
			}
		case 6:
			//JUMP if false
			//	fmt.Println(b[i], b[i+1], b[i+2])

			first := getParamterOutput(modeC, relativeBase, i, b, 1)

			second := getParamterOutput(modeB, relativeBase, i, b, 2)

			if b[first] == 0 {
				i = b[second]
			} else {
				i = i + 3
			}
		case 7:
			//fmt.Println(b[i], b[i+1], b[i+2], b[i+3])

			first := getParamterOutput(modeC, relativeBase, i, b, 1)

			second := getParamterOutput(modeB, relativeBase, i, b, 2)
			third := getParamterOutput(modeA, relativeBase, i, b, 3)

			if b[first] < b[second] {
				b[third] = 1
			} else {
				b[third] = 0
			}
			i = i + 4

		case 8:
			//			fmt.Println(b[i], b[i+1], b[i+2], b[i+3])

			first := getParamterOutput(modeC, relativeBase, i, b, 1)

			second := getParamterOutput(modeB, relativeBase, i, b, 2)
			third := getParamterOutput(modeA, relativeBase, i, b, 3)

			if b[first] == b[second] {
				b[third] = 1
			} else {
				b[third] = 0
			}
			i = i + 4
		case 9:
			//first := getParamterOutput(modeC, relativeBase, i, b, 1)
			switch modeC {
			case 0:
				relativeBase = relativeBase + b[b[i+1]]
			case 1:
				relativeBase = relativeBase + b[i+1]

			case 2:
				relativeBase = relativeBase + b[relativeBase+b[i+1]]

			}
			i = i + 2
		case 99:
			if name == "COMP1" {
				firstOneRunning = false
			}
			fmt.Println("Painted,", painted)
			for i := 0; i < 150; i++ {
				for j := 0; j < 150; j++ {
					if shipHull[i][j] == 0 {
						fmt.Print(" ")
					} else {
						fmt.Print("#")
					}

				}
				fmt.Println("")
			}
			return outputParam
		default:
			fmt.Println("Unknown,", op)
			i = i + 4
			return 3
		}

	}
	return outputParam
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func getParamterOutput(modParam int, relativeBase int, index int, program []int, pos int) (resultingAdress int) {
	switch modParam {
	case 0:
		//fmt.Println("0, value of value ")
		resultingAdress = program[index+pos]

	case 1:
		//fmt.Println("1, stay the same")
		resultingAdress = index + pos
	case 2:
		resultingAdress = relativeBase + program[index+pos]
	default:
		fmt.Println("ERROR", program[index], program[index+1], program[index+2], program[index+3])
	}
	return
}
