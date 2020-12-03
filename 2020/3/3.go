package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var row int
	terrain := make(map[string]string)
	for scanner.Scan() {

		line := scanner.Text()
		for i := range line {
			key := strconv.Itoa(row) + "-" + strconv.Itoa(i)
			fmt.Println(row, i)
			if string(line[i]) == "#" {
				terrain[key] = "tree"
			} else {
				terrain[key] = "open"
			}
		}
		row = row + 1
	}

	fmt.Println(cruise(terrain, 1, 1) *
		cruise(terrain, 3, 1) *
		cruise(terrain, 5, 1) *
		cruise(terrain, 7, 1) *
		cruise(terrain, 1, 2))
}

func cruise(terrain map[string]string, right int, down int) (trees int) {
	a := 0
	width := 31
	trees = 0
	level := 0
	for level < 1000 {

		key := strconv.Itoa(level) + "-" + strconv.Itoa(a)
		a = (a + right) % width
		if terrain[key] == "tree" {
			trees = trees + 1
		}
		level = level + down
	}
	fmt.Println("Nbr of Trees", trees)
	return trees
}
