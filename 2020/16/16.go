package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	name     string
	lowMin   int
	lowMax   int
	upperMin int
	upperMax int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var rules []Rule
	allRules := make([][]Rule, 20)
	okeyEdRules := make([]Rule, 20)
	parseRules := true
	parseMyTick := false
	parseNearby := false
	nonValidSum := 0
	var myTicket []int
	for scanner.Scan() {
		line := scanner.Text()
		if parseRules {
			if line == "" {
				parseRules = false
				parseMyTick = true
			} else {
				rule := strings.Split(line, ":")
				bounds := strings.Split(rule[1], " ")
				lowers := strings.Split(bounds[1], "-")
				uppers := strings.Split(bounds[3], "-")

				lMin, _ := strconv.Atoi(lowers[0])
				lMax, _ := strconv.Atoi(lowers[1])
				uMin, _ := strconv.Atoi(uppers[0])
				uMax, _ := strconv.Atoi(uppers[1])

				r := Rule{name: rule[0], lowMin: lMin, lowMax: lMax, upperMin: uMin, upperMax: uMax}
				rules = append(rules, r)

			}
		} else if parseMyTick {

			if line != "your ticket:" {
				if line == "" {
					parseMyTick = false
					parseNearby = true
				} else {
					//Create PotRuleMap
					numFields := strings.Split(line, ",")
					for i := 0; i < len(numFields); i++ {
						allRules[i] = append(allRules[i], rules...)
					}

					for _, a := range numFields {
						k, _ := strconv.Atoi(a)
						myTicket = append(myTicket, k)
					}

					valid, nonValids := validateTicket(line, rules)
					for _, nv := range nonValids {
						nonValidSum = nonValidSum + nv
					}

					allRules = excludeRules(line, allRules)

					if valid {

					}
				}
			}
		} else if parseNearby {
			if line != "nearby tickets:" {
				if line == "" {

				} else {
					valid, nonValids := validateTicket(line, rules)

					for _, nv := range nonValids {
						nonValidSum = nonValidSum + nv
					}
					if valid {
						allRules = excludeRules(line, allRules)

					}
				}
			}
		}
	}
	fmt.Println("nonvalids", nonValidSum)

	z := 0
	for {

		for i, r := range allRules {

			if len(r) == 1 {
				okeyEdRules[i] = r[0]
				allRules[i] = []Rule{}
				for y, w := range allRules {
					foundRule := -1
					for j, u := range w {
						if u.name == r[0].name {
							foundRule = j
						}
					}
					if foundRule != -1 {
						allRules[y] = RemoveIndex(w, foundRule)
					}
				}
			}
		}
		z++
		if z > 20 {
			break
		}

	}
	depSum := 1
	for i, r := range okeyEdRules {

		fmt.Println(i, r)
		if strings.HasPrefix(r.name, "departure") {
			depSum = depSum * myTicket[i]
		}
	}
	fmt.Println("Sum2 ", depSum)

}

func validateTicket(nums string, rules []Rule) (valid bool, nonValids []int) {

	numbers := strings.Split(nums, ",")
	for _, n := range numbers {
		v := false
		a, _ := strconv.Atoi(n)
		for _, r := range rules {
			if (r.lowMin <= a && r.lowMax >= a) || (r.upperMin <= a && r.upperMax >= a) {
				v = true
			}
		}
		if !v {
			nonValids = append(nonValids, a)
			valid = false
			break
		}
		valid = true
	}
	return
}

func excludeRules(nums string, allRules [][]Rule) (result [][]Rule) {
	numbers := strings.Split(nums, ",")
	for i, n := range numbers {

		var removeTheseRules []Rule
		a, _ := strconv.Atoi(n)
		for _, r := range allRules[i] {
			if (r.lowMin <= a && r.lowMax >= a) || (r.upperMin <= a && r.upperMax >= a) {
			} else {
				removeTheseRules = append(removeTheseRules, r)

			}
		}
		for _, rr := range removeTheseRules {
			indexToRemove := -1
			for j, r := range allRules[i] {
				if rr.name == r.name {
					indexToRemove = j
				}
			}
			if indexToRemove != -1 {
				allRules[i] = RemoveIndex(allRules[i], indexToRemove)

			}
		}

	}
	return allRules
}

func RemoveIndex(s []Rule, index int) []Rule {
	return append(s[:index], s[index+1:]...)
}
