package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	a := strings.Split(line, ",")
	orig := make([]int, len(a))
	for i, v := range a {
		orig[i], _ = strconv.Atoi(v)

	}
	fmt.Println(orig[0])
	//1202 program alarm

	noun := 0
	for noun < 99 {

		verb := 0
		for verb < 99 {
			b := append(orig[:0:0], orig...)
			//b[1] = noun
			//b[2] = verb
			res := run(b)
			//fmt.Println(noun, verb, res)
			if res == 19690720 {
				//	fmt.Println(noun, verb)
				break
			}
			verb++
		}
		noun++
	}

}

/*
Opcode 3 takes a single integer as input and saves it to the position given by its only parameter. For example, the instruction 3,50 would take an input value and store it at address 50.
Opcode 4 outputs the value of its only parameter. For example, the instruction 4,50 would output the value at address 50.
*/
func run(b []int) int {
	inputParam := 5
	outputParam := 0
	i := 0

	for i < len(b) {

		//fmt.Println("OP", b[i]%10)
		op := b[i] % 10
		//modeA := b[i] / 10000
		modeB := b[i] / 1000
		modeC := b[i] / 100 % 10
		switch op {
		case 01:

			//add
			firstParam := b[i+1]
			if modeC == 0 {
				firstParam = b[b[i+1]]

			}
			secondParam := b[i+2]

			if modeB == 0 {
				secondParam = b[b[i+2]]

			}
			b[b[i+3]] = firstParam + secondParam
			i = i + 4

		case 02:
			//Multiply
			//	fmt.Println("Multiply")
			firstParam := b[i+1]
			if modeC == 0 {
				firstParam = b[b[i+1]]

			}
			secondParam := b[i+2]

			if modeB == 0 {
				secondParam = b[b[i+2]]

			}

			b[b[i+3]] = firstParam * secondParam
			i = i + 4
		case 03:
			b[b[i+1]] = inputParam
			i = i + 2

		case 04:
			outputParam = b[b[i+1]]
			i = i + 2
			if outputParam != 0 {
				fmt.Println(outputParam)

			}

		case 05:
			//JUMP if true
			firstParam := b[i+1]

			if modeC == 0 {
				firstParam = b[b[i+1]]
			}
			secondParam := b[i+2]

			if modeB == 0 {
				secondParam = b[b[i+2]]

			}

			if firstParam != 0 {
				i = secondParam
			} else {
				i = i + 3
			}
		case 06:
			//JUMP if false
			firstParam := b[i+1]
			if modeC == 0 {
				firstParam = b[b[i+1]]
			}
			secondParam := b[i+2]

			if modeB == 0 {
				secondParam = b[b[i+2]]

			}

			if firstParam == 0 {
				i = secondParam
			} else {
				i = i + 3
			}
		case 07:
			//Opcode 7 is less than: if the first parameter is less than the second parameter,
			//it stores 1 in the position given by the third parameter. Otherwise, it stores 0.

			firstParam := b[i+1]
			if modeC == 0 {
				firstParam = b[b[i+1]]

			}
			secondParam := b[i+2]

			if modeB == 0 {
				secondParam = b[b[i+2]]

			}
			if firstParam < secondParam {
				b[b[i+3]] = 1
			} else {
				b[b[i+3]] = 0
			}

			i = i + 4

		case 8:
			//	Opcode 8 is equals: if the first parameter is equal to the second parameter,
			//	it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			firstParam := b[i+1]
			if modeC == 0 {
				firstParam = b[b[i+1]]

			}
			secondParam := b[i+2]

			if modeB == 0 {
				secondParam = b[b[i+2]]

			}
			if firstParam == secondParam {
				b[b[i+3]] = 1
			} else {
				b[b[i+3]] = 0
			}

			i = i + 4
		case 9:
			return b[0]
		default:
			fmt.Println("Unknown,", b[i]%10)
			i = i + 4
		}

	}
	//fmt.Println(outputParam)

	return b[0]
}
