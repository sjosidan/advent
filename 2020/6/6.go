package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	questions := make([]rune, 0)
	var totals int
	newTeam := true
	for scanner.Scan() {

		line := scanner.Text()
		if line == "" {
			totals = totals + len(questions)
			questions = make([]rune, 0)
			newTeam = true
		} else if newTeam {
			for _, v := range line {
				questions = append(questions, v)
			}
			newTeam = false
		} else {
			refinedQuestions := make([]rune, 0)
			for _, v := range questions {
				if strings.Contains(line, string(v)) {
					refinedQuestions = append(refinedQuestions, v)
				}
			}
			questions = refinedQuestions
		}

	}
	//Add last team
	totals = totals + len(questions)
	//Print totals
	fmt.Println("total questions", totals)
}
