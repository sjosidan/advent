package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Seat struct {
	row    int64
	column int64
	id     int64
}

/*
* This solution indexes all the Boarding passes and fills the map with taken seats.
 */

func main() {
	boardingPasses := make(map[int64]Seat)
	// READ DATA
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		partOne := strings.Replace(strings.Replace(line[:7], "F", "0", 10), "B", "1", 10)
		partTwo := strings.Replace(strings.Replace(line[7:len(line)], "L", "0", 10), "R", "1", 10)
		roww, _ := strconv.ParseInt(partOne, 2, 64)
		coll, _ := strconv.ParseInt(partTwo, 2, 64)

		boardingPasses[(roww*8)+(coll)] = Seat{row: roww, column: coll, id: (roww * 8) + (coll)}
	}

	//SORT AND PRINT
	keys := make([]int64, 0, len(boardingPasses))
	for k := range boardingPasses {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for i, k := range keys {
		if i != 0 && (keys[i-1]-k != -1) {
			fmt.Println("YOUR SEAT", k-1)
			break
		}

	}

}
