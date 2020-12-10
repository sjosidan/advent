package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	statements := make([]string, 0)
	//masterSlice := make([]string, 0)
	for scanner.Scan() {

		line := scanner.Text()
		statements = append(statements, line)
	}

	result := getSlice("shiny gold", statements)
	ms := uniqueNonEmptyElementsOf(result)
	fmt.Println(len(ms)-1, ms)
}

func getSlice(seed string, buf []string) (result []string) {
	fmt.Println("Looking for ", seed)
	for _, a := range buf {
		b := strings.Split(a, " ")
		if strings.Contains(a, seed) && b[0]+" "+b[1] != seed {
			fmt.Println("Found match ", b[0]+" "+b[1], "for seed", seed)
			//result = append(result, b[0]+" "+b[1])
			result = append(result, append(getSlice(b[0]+" "+b[1], buf))...)
		}
		if b[0]+" "+b[1] == seed {
			result = append(result, b[0]+" "+b[1])
		}
	}

	return
}

func uniqueNonEmptyElementsOf(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}

	return us

}
