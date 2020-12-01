package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	j := 0

	s := make([]int, 200)
	for scanner.Scan() {
		line := scanner.Text()
		i, _ := strconv.Atoi(line)

		s[j] = i
		j = j + 1
	}
	fmt.Println(s)

	for k := 0; k < len(s); k++ {
		for _, e1 := range s {
			for _, e2 := range s {
				if s[k]+e1+e2 == 2020 {
					fmt.Println("2020 ", s[k], e1, e2, "answ ", s[k]*e1*e2)
				}
			}
		}
	}
}
