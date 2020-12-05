package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Seat struct {
	row    int
	column int
	id     int
}

func main() {
	rowRange := makeRange(0, 127)
	colRange := makeRange(0, 7)
	var searchRange []int
	airPlane := make(map[string]Seat)

	//BUILD AN AIRPLANE
	for i := 0; i < 128; i++ {
		searchRange = rowRange

		var rowString string
		for len(searchRange) != 1 {
			lowerSide := searchRange[0 : len(searchRange)/2]
			upperSide := searchRange[len(searchRange)/2 : len(searchRange)]
			resultLower := sort.SearchInts(lowerSide, i)
			if resultLower == len(lowerSide) {
				searchRange = upperSide
				rowString = rowString + "B"
			} else {
				searchRange = lowerSide
				rowString = rowString + "F"
			}
		}

		searchRange = colRange

		for j := 0; j < 8; j++ {
			searchRange = colRange
			var columnString string
			for len(searchRange) != 1 {
				lowerSide := searchRange[0 : len(searchRange)/2]
				upperSide := searchRange[len(searchRange)/2 : len(searchRange)]
				resultLower := sort.SearchInts(lowerSide, j)
				if resultLower == len(lowerSide) {
					searchRange = upperSide
					columnString = columnString + "R"
				} else {
					searchRange = lowerSide
					columnString = columnString + "L"
				}
			}
			totti := rowString + columnString
			airPlane[totti] = Seat{row: i, column: j, id: (i * 8) + (j)}
		}
	}

	// READ DATA
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		//DELTE Taken seats
		delete(airPlane, line)
	}

	// Print the Remaning seats
	/*
		for a, v := range airPlane {
			fmt.Println(a, v)
		}
	*/

	//SORT AND PRINT
	keys := make([]string, 0, len(airPlane))

	for k := range airPlane {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, airPlane[k])
	}

}

func makeRange(min int, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
