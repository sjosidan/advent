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
			b[1] = noun
			b[2] = verb
			res := run(b)
			fmt.Println(noun, verb, res)
			if res == 19690720 {
				fmt.Println(noun, verb)
				break
			}
			verb++
		}
		noun++
	}

}
func run(b []int) int {
	i := 0

	for i < len(b) {

		switch b[i] {
		case 1:
			//add

			b[b[i+3]] = b[b[i+1]] + b[b[i+2]]
		case 2:
			//Multiply
			//fmt.Println("mult", b[i+1], b[i+2], b[i+3])
			b[b[i+3]] = b[b[i+1]] * b[b[i+2]]

		case 99:
			return b[0]
		}
		i = i + 4
	}
	return b[0]
}
