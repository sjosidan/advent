package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var questions string
	var totals int
	newTeam := true
	for scanner.Scan() {

		line := scanner.Text()
		if line == "" {
			totals = totals + len(questions)
			newTeam = true
		} else if newTeam {
			questions = line
			newTeam = false
		} else {
			refinedQuestions := make([]rune, 0)
			for _, v := range questions {
				if strings.Contains(line, string(v)) {
					refinedQuestions = append(refinedQuestions, v)
				}
			}
			questions = string(refinedQuestions)
		}

	}
	//Add last team
	totals = totals + len(questions)
	//Print totals
	fmt.Println("total questions", totals)
}
