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
	row    int
	column int
	id     int
}

func main() {
	airPlane := make(map[string]Seat)

	//BUILD AN AIRPLANE
	for i := 0; i < 128; i++ {
		l := strconv.FormatInt(int64(i), 2)
		switch len(l) {
		case 1:
			l = "000000" + l
		case 2:
			l = "00000" + l
		case 3:
			l = "0000" + l
		case 4:
			l = "000" + l
		case 5:
			l = "00" + l
		case 6:
			l = "0" + l
		}
		row := strings.Replace(strings.Replace(l, "0", "F", 10), "1", "B", 10)

		for j := 0; j < 8; j++ {
			l := strconv.FormatInt(int64(j), 2)
			switch len(l) {
			case 1:
				l = "00" + l
			case 2:
				l = "0" + l
			}
			column := strings.Replace(strings.Replace(l, "0", "L", 10), "1", "R", 10)
			airPlane[row+column] = Seat{row: i, column: j, id: (i * 8) + (j)}
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
