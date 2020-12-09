package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	preAmbleLength := 25
	var stream []int
	for scanner.Scan() {
		line := scanner.Text()
		v, _ := strconv.Atoi(line)

		stream = append(stream, v)
	}
	j := preAmbleLength
	for i := preAmbleLength; i < len(stream); i++ {
		for j = range stream[i-preAmbleLength : i] {
			match := false
			for k := range stream[i-preAmbleLength : i] {
				if stream[i-preAmbleLength+j]+stream[i-preAmbleLength+k] == stream[i] {
					match = true
				}
			}
			if match {
				break
			}
			if j == len(stream[i-preAmbleLength:i])-1 {
				fmt.Println("no-match ", stream[i])
				findMatch(stream[i], stream)
			}
		}
	}

}

func findMatch(findMe int, stream []int) {
	match := false
	i := 0
	for !match {

		for i < len(stream) {
			result := 0
			for k, v := range stream[i:len(stream)] {
				result += v
				if result > findMe {
					break
				}
				if result == findMe {
					fmt.Println("found it ", stream[i:i+k])
					min, max := minMax(stream[i : i+k])
					fmt.Println("min", min, "max", max)
					fmt.Println(min + max)
					return
				}
			}
			i++
		}

	}

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
