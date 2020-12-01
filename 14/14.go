package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		conv := strings.Split(line, "=>")
		for i := 0; i < len(conv)-1; i++ {

			for _, a := range strings.Split(conv[i], ",") {
				fmt.Println(a)
			}
		}
		fmt.Println("Converts to", conv[len(conv)-1])

	}
}
