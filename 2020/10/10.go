package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var stream []int
	stream = append(stream, 0)
	for scanner.Scan() {
		line := scanner.Text()
		v, _ := strconv.Atoi(line)
		stream = append(stream, v)
	}

	_, max := minMax(stream)
	stream = append(stream, max+3)
	sort.Ints(stream)
	fmt.Println(stream)
	oneVolters := 0
	threeVolters := 0

	for i, z := range stream {
		if i != 0 {
			if z-stream[i-1] == 3 {
				threeVolters++
			} else if z-stream[i-1] == 1 {
				oneVolters++
			}

		}
	}
	fmt.Println("onevolt\t", oneVolters)
	fmt.Println("threevolt\t", threeVolters)
	isNeeded(stream)
}

func minMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func isNeeded(myNiceSlice []int) (result int) {
	result = 0

	for i, a := range myNiceSlice {
		if (i != 0) && (i < len(myNiceSlice)-2) {
			if myNiceSlice[i+1]-myNiceSlice[i-1] < 4 {
				//copycat := myNiceSlice
				copycat := append(myNiceSlice[:i], myNiceSlice[i+1:]...)
				fmt.Println(a, "is not needed", copycat)
				result = 1 + isNeeded(copycat)
			}
		}
	}

	return
}
