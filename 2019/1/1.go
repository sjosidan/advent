package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	acc := 0
	for scanner.Scan() {
		line := scanner.Text()
		i, _ := strconv.Atoi(line)
		res := (i / 3) - 2
		acc = acc + res
		// this is for second input
		/*
			for res > 0 && res > 3 {
				res = (res / 3) - 2
				if res > 0 {
					acc = acc + res
				}

			}
		*/
	}
	fmt.Println(acc)

}
