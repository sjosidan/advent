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
	onediff := true
	firstIndex := 1
	tot := 1
	for i, _ := range stream {
		if i > 1 && (len(stream)) > i {
			if stream[i]-stream[i-1] == 1 {
				if !onediff {
					firstIndex = i
				}
				onediff = true
			} else {
				if onediff {
					switch len(stream[firstIndex-1 : i]) {
					case 3:
						tot = tot * 2
					case 4:
						tot = tot * 4
					case 5:
						tot = tot * 7
					default:
						fmt.Println("NO VAL LEN ", len(stream[firstIndex-1:i]))
					}
				}
				onediff = false
			}
		}
	}
	fmt.Println(tot)
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
