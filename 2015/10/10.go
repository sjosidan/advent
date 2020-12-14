package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Ones struct {
	index  int
	ones   string
	length int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < 2; i++ {

			line = lookAndSay(line)
		}
		fmt.Println(len(line))
	}
}

func lookAndSay(input string) (result string) {

	c := input[0]
	nc := 1
	for i := 1; i < len(input); i++ {
		d := input[i]
		if d == c {
			nc++
			continue
		}
		result += strconv.Itoa(nc) + string(c)
		c = d
		nc = 1
	}
	return result + strconv.Itoa(nc) + string(c)

}
