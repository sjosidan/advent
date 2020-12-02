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
	var valid1 int
	var valid2 int
	var nonvalid2 int
	for scanner.Scan() {
		line := scanner.Text()

		a := strings.Split(line, " ")
		bounds := strings.Split(a[0], "-")
		min, _ := strconv.Atoi(bounds[0])
		max, _ := strconv.Atoi(bounds[1])

		key := a[1][0:1]
		res := strings.Count(a[2], key)
		if res >= min && res <= max {
			valid1 = valid1 + 1
		} else {
		}

		firstPos := a[2][min-1:min] == key
		secondPos := a[2][max-1:max] == key

		if (firstPos || secondPos) && !(firstPos && secondPos) {
			valid2 = valid2 + 1
		} else {
			nonvalid2 = nonvalid2 + 1

		}

	}
	fmt.Println("Valid", valid2)
	fmt.Println("NonValid", nonvalid2)

}
