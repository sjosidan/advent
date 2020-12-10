package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Contents struct {
	bagColor string
	weight   string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	statements := make([]string, 0)
	bagMap := make(map[string][]Contents)
	for scanner.Scan() {

		line := scanner.Text()
		statements = append(statements, line)
		manyBags := strings.Split(line, " ")
		contents := make([]Contents, 0)
		for j := 1; j <= len(manyBags)-5; {
			contents = append(contents, Contents{bagColor: manyBags[j+4] + " " + manyBags[j+5], weight: manyBags[j+3]})
			j = j + 4
		}
		bagMap[manyBags[0]+" "+manyBags[1]] = contents
	}

	tot := count("shiny gold", bagMap)
	fmt.Println(tot)
}

func count(seed string, bag map[string][]Contents) (tot int) {
	for _, a := range bag[seed] {
		fmt.Println("Adding ", a.bagColor, a.weight)
		weight, _ := strconv.Atoi(a.weight)
		tot = tot + weight
		tot = tot + weight*count(a.bagColor, bag)
	}
	return
}
