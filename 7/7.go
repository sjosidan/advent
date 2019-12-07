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
	arr := []int{5, 6, 7, 8, 9}
	for _, a := range permutations(arr) {
		wg := new(sync.WaitGroup)

		input1 := make(chan int)
		input2 := make(chan int, a[1])
		input3 := make(chan int, a[2])
		input4 := make(chan int, a[3])
		input5 := make(chan int, a[4])

		b := append(orig[:0:0], orig...)
		wg.Add(1)

		go run(b, &input1, &input2, "COMP1", wg)

		b = append(orig[:0:0], orig...)
		wg.Add(1)

		go run(b, &input2, &input3, "COMP2", wg)

		b = append(orig[:0:0], orig...)
		wg.Add(1)
		go run(b, &input3, &input4, "COMP3", wg)

		b = append(orig[:0:0], orig...)
		wg.Add(1)
		go run(b, &input4, &input5, "COMP4", wg)

		b = append(orig[:0:0], orig...)
		wg.Add(1)
		go run(b, &input5, &input1, "COMP5", wg)
		input1 <- a[0]
		input2 <- a[1]
		input3 <- a[2]
		input4 <- a[3]
		input5 <- a[4]
		input1 <- 0

		fmt.Println(a)

		wg.Wait()

		close(input1)
		close(input2)
		close(input3)
		close(input4)
		close(input5)

	}

	fmt.Println("MAIN Sleeping")
	fmt.Println(globalMax)

}

func run(b []int, inputChannel *chan int, outputChannel *chan int, name string, wg *sync.WaitGroup) int {

	defer wg.Done()
	if name == "COMP1" {
		firstOneRunning = true
	}
	outputParam := 0
	i := 0

	for i < len(b) {

		op := b[i] % 10
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
			inputStuff := <-*inputChannel
			b[b[i+1]] = inputStuff
			i = i + 2

		case 04:
			outputParam = b[b[i+1]]
			if name == "COMP5" {
				if !firstOneRunning {
					fmt.Println(name, "Done with last one", b[b[i+1]])
					if globalMax < b[b[i+1]] {
						globalMax = b[b[i+1]]
					}
				} else {
					*outputChannel <- b[b[i+1]]
				}

			} else {
				*outputChannel <- b[b[i+1]]
			}
			i = i + 2

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
			if name == "COMP1" {
				firstOneRunning = false
			}
			return outputParam
		default:
			fmt.Println("Unknown,", b[i]%10)
			i = i + 4
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
