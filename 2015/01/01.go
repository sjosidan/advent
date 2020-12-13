package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	floor := 0
	for scanner.Scan() {
		line := scanner.Text()
		for index, l := range line {
			switch string(l) {
			case "(":
				floor++
			case ")":
				floor--
			default:

			}
			if floor == -1 {
				fmt.Println("Entered basement at", index+1)
			}
		}
	}
	fmt.Println("Floor:", floor)
}
