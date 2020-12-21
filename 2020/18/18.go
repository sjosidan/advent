package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	start int
	end   int
}

type key struct {
	key string
	val int
}

type ByLen []string

func (a ByLen) Len() int           { return len(a) }
func (a ByLen) Less(i, j int) bool { return len(a[i]) > len(a[j]) }
func (a ByLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	tot := 0

	for scanner.Scan() {

		line := scanner.Text()
		openParants := 0
		var lastOpenIndexes []int
		parseString := line
		z := 0
		for {
			newString := ""
			if strings.Contains(parseString, "(") {
				for i, l := range parseString {

					if string(l) == "(" {
						openParants++
						lastOpenIndexes = append(lastOpenIndexes, i)
					}
					if string(l) == ")" {
						openParants--
						lastIndex := lastOpenIndexes[len(lastOpenIndexes)-1]
						replaceMe := parseString[lastIndex : i+1]
						//fmt.Println(replaceMe)
						ans := calc(replaceMe[1 : len(replaceMe)-1])
						newString = strings.Replace(parseString, replaceMe, strconv.Itoa(ans), -1)
						break
					}
				}
			} else {
				//fmt.Println("done", line, calc(parseString))
				tot += calc(parseString)
				break
			}
			parseString = newString
			z++
			//fmt.Println(parseString)
			if z == 15 {
				fmt.Println("WFTC", parseString)
				break
			}
		}

	}
	fmt.Println("alles", tot)
	//fmt.Println(calc("2 * 5"))

	/*
		for scanner.Scan() {
			var leftOpens []int
			var parat []Pair
			reps := make(map[string]int)
			k := ""
			line := scanner.Text()
			k = line
			leftOpen := 0

			for a, l := range line {
				if string(l) == "(" {
					leftOpen++
					leftOpens = append(leftOpens, a)
				}
				if string(l) == ")" {
					leftOpen--
					p := Pair{start: leftOpens[len(leftOpens)-1], end: a + 1}
					leftOpens = leftOpens[:len(leftOpens)-1]
					parat = append(parat, p)
				}

			}
			var keys []key
			for _, a := range parat {
				z := evalExpr(k[a.start:a.end])
				reps[k[a.start:a.end]] = z
				keys = append(keys, key{key: k[a.start:a.end], val: z})
			}
			z := 0

			for {
				for _, i := range keys {
					tempKey := i.key

					for _, r := range keys {
						if i.key != r.key {
							if strings.Contains(tempKey, r.key) {
								tempKey = strings.Replace(tempKey, r.key, strconv.Itoa(r.val), -1)
								nn := evalExpr(tempKey)
								reps[i.key] = nn
								//	fmt.Println(tempKey, r.key, r.val)
								delete(reps, r.key)

							}
						}
					}
				}
				z++
				//fmt.Println(z)
				if z > 20 {
					break
				}
			}
			//fmt.Println(reps)

			a := line
			for i, k := range reps {
				a = strings.Replace(a, i, strconv.Itoa(k), -1)
			}
			loc := evalExpr(a)
			tot = tot + loc
			fmt.Println(a, loc)
		}
	*/

}

func calc(a string) (result int) {
	//	fmt.Println("input", a)
	s := strings.Split(a, " ")

	stillPlus := false
	additionMap := make(map[string]int)
	var additionSlice []int
	additionKey := ""
	//ADD EVERYTHING
	for i := 1; i < len(s); i++ {
		if s[i] == "+" {
			n, _ := strconv.Atoi(s[i-1])
			stillPlus = true
			additionSlice = append(additionSlice, n)
			additionKey = additionKey + s[i-1] + " " + s[i] + " "
			if i == len(s)-2 {
				tot := 0
				n, _ := strconv.Atoi(s[i+1])

				additionKey = additionKey + s[i+1]
				additionSlice = append(additionSlice, n)

				for _, a := range additionSlice {
					tot += a
				}
				additionMap[additionKey] = tot
				stillPlus = false
			}
		} else {
			if stillPlus {
				tot := 0
				n, _ := strconv.Atoi(s[i-1])
				additionSlice = append(additionSlice, n)
				additionKey = additionKey + s[i-1]
				for _, a := range additionSlice {
					tot += a
				}
				additionMap[additionKey] = tot
				stillPlus = false
				additionKey = ""
				additionSlice = []int{}

			}
		}
		i++
	}
	//fmt.Println(additionMap)
	//MULTIPLY WHATS REST
	//Sort keys by length
	var addKeys []string
	for k := range additionMap {
		addKeys = append(addKeys, k)
	}
	sort.Sort(ByLen(addKeys))
	cc := a
	for _, k := range addKeys {
		val := strconv.Itoa(additionMap[k])
		//	fmt.Println("replace", k, "with", val)
		multiString := strings.Replace(cc, k, val, -1)
		cc = multiString
	}
	//fmt.Println(cc)
	tot := 1
	for _, l := range strings.Split(cc, " ") {
		if l != "*" {
			val, _ := strconv.Atoi(l)
			tot *= val
		}
	}
	result = tot
	//fmt.Println("output", tot)
	return
}

func evalExpr(line string) (tot int) {
	s := strings.Split(line, " ")
	tot, _ = strconv.Atoi(strings.Replace(s[0], "(", "", -1))
	var unescp []string
	for _, z := range s {
		n := strings.Replace(z, ")", "", -1)
		n = strings.Replace(n, "(", "", -1)
		unescp = append(unescp, n)
	}
	for i := 1; i < len(unescp)-1; i++ {

		if s[i] == "+" {
			n, err := strconv.Atoi(unescp[i+1])
			tot = tot + n
			if err != nil {
				fmt.Println(err)
			}
		} else if s[i] == "*" {
			n, err := strconv.Atoi(unescp[i+1])
			tot = tot * n
			if err != nil {
				fmt.Println(err)
			}

		} else {
			fmt.Println("WTFF", line)
		}
		i++

	}
	//fmt.Println(line, " == ", tot)

	return
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
