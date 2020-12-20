package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Rule struct {
	id string
	s  string
}

var ruleMap = make(map[string]Rule)
var oKRules = make(map[string]bool)
var oKRules2 = make(map[string]bool)
var trimRules = make(map[string]bool)
var trimRules2 = make(map[string]bool)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	parseRules := true
	var solvedRules []Rule
	var unsolvedRules []Rule
	var rules []Rule
	buildRule := true
	matches := 0
	for scanner.Scan() {
		line := scanner.Text()
		match := false
		kk := strings.Split(line, " ")

		if line == "" {
			parseRules = false
		} else if parseRules {
			//rLeng := len(kk[0]) + 1
			//	if !(strings.HasPrefix(line, "8:")) {
			//		if !(strings.HasPrefix(line, "11:")) {
			r := Rule{id: kk[0][:len(kk[0])-1], s: line[len(kk[0])+1:]}
			if kk[1] == "\"b\"" {
				r = Rule{id: kk[0][:len(kk[0])-1], s: "b"}
				solvedRules = append(solvedRules, r)

			} else if kk[1] == "\"a\"" {
				r = Rule{id: kk[0][:len(kk[0])-1], s: "a"}
				solvedRules = append(solvedRules, r)

			} else {
				unsolvedRules = append(unsolvedRules, r)
			}
			rules = append(rules, r)
			ruleMap[r.id] = r
			//		}
			//	}
		} else {
			if buildRule {
				fmt.Println(ruleMap["42"].s)
				fmt.Println(ruleMap["31"].s)
				expandRule(ruleMap["888"].s, rules, 1)
				expandRule(ruleMap["999"].s, rules, 2)

				for a := range oKRules {
					trimmed := strings.Replace(a, " ", "", -1)
					trimRules[trimmed] = true
					fmt.Println("042", len(trimmed), trimmed)
				}

				for a := range oKRules2 {
					trimmed2 := strings.Replace(a, " ", "", -1)
					trimRules2[trimmed2] = true
					fmt.Println("031", len(trimmed2), trimmed2)
				}

				buildRule = false
			}
			if trimRules[line] {
				//				matches++
			}

			left := line
			leftOK := false

			u := 0
			p := 0
			q := 0

			for {
				for l := range trimRules {

					if strings.HasPrefix(left, l) {
						left = line[(u+1)*len(l):]
						leftOK = true
						u++
						break
					}
				}
				p++
				if p == 20 {
					break
				}
			}
			numOfRights := 0
			if leftOK {
				for {
					for r := range trimRules2 {
						if strings.HasPrefix(left, r) {

							numOfRights++
							if left == r && ((u + 1 - numOfRights) > numOfRights) {
								matches++
								match = true
								q = 19
								break
							} else if left == r {
								fmt.Println("doesnt full exp, left", u-numOfRights, numOfRights)
							}
							left = line[(u+1)*len(r):]
							u++

						}
					}
					q++
					if q == 20 {
						break
					}
				}
			}
		}
		if !match {
			fmt.Println("NO MATCH", line, len(line))
		}

	}

	fmt.Println("Matches", matches)

}

func expandRule(r string, listOfRules []Rule, iteration int) {
	//var ll []string
	z := ""
	spliiter := strings.Split(r, " ")
	//fmt.Println("----", spliiter)
	for i, b := range spliiter {
		if strings.Contains(ruleMap[b].s, "|") {

			z = z + " (" + ruleMap[b].s + ") "
		} else {
			if i != 0 {
				z = z + " " + ruleMap[b].s

			} else {
				z = z + ruleMap[b].s + " "
			}
			if b == "b" || b == "a" {
				z = z + b
			}

		}

	}
	z = strings.Replace(z, "  ", " ", -1)
	// fmt.Println(z)

	k := z
	x := 0
	expandMe := make(map[string]bool)
	//fmt.Println(len(expandMe))
	var pooo map[string]bool
	firstTime := true
	for {
		if firstTime {
			if strings.Contains(k, "|") {
				p := strings.Index(k, "|")
				l := strings.Index(k, "(")
				r := strings.Index(k, ")")

				lef := (k[l : p+1])
				rig := (k[p : r+1])
				a := strings.Replace(strings.Replace(k, lef, "", 1), ")", "", 1)
				b := strings.Replace(strings.Replace(k, rig, "", 1), "(", "", 1)
				a = strings.Replace(a, "  ", " ", -1)
				b = strings.Replace(b, "  ", " ", -1)

				if string(a[0]) == " " {
					a = a[1:]
				}
				if string(b[0]) == " " {
					b = b[1:]
				}
				expandMe[a] = true
				expandMe[b] = true

			} else {
				expandMe[k] = true
			}

			firstTime = false
		} else {
			for k = range pooo {
				if strings.Contains(k, "|") {
					p := strings.Index(k, "|")
					l := strings.Index(k, "(")
					r := strings.Index(k, ")")
					lef := (k[l : p+1])
					rig := (k[p : r+1])
					a := strings.Replace(strings.Replace(k, lef, "", 1), ")", "", 1)
					b := strings.Replace(strings.Replace(k, rig, "", 1), "(", "", 1)

					a = strings.Replace(a, "  ", " ", -1)
					b = strings.Replace(b, "  ", " ", -1)
					if string(a[0]) == " " {
						a = a[1:]
					}
					if string(b[0]) == " " {
						b = b[1:]
					}

					expandMe[a] = true
					expandMe[b] = true

				} else {
				}
			}
			pooo = expandMe
		}
		x++
		if x == 15 {
			break
		}

	}

	for i := range expandMe {
		//	fmt.Println(i)
		if alphaOnly(i) {
			if iteration == 1 {
				//fmt.Println("adding to 042")
				oKRules[i] = true
			} else {
				//	fmt.Println("adding to 031")
				oKRules2[i] = true
			}
		} else {
			//fmt.Println(len(expandMe))
			//fmt.Println(i)
			if !strings.Contains(i, "|") {
				expandRule(i, listOfRules, iteration)
			}
		}
	}
	// fmt.Println(expandMe)
	/*ll = append(ll, z)
	for r, z := range ll {

	}
	*/
	//fmt.Println("rule ", z, " ")
	//	fmt.Println(ll)

}

const alpha = "ab "

func alphaOnly(s string) bool {
	for _, char := range s {
		if !strings.Contains(alpha, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}
