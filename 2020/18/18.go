package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	tot := 0
	/*
		var plussings []int

		var q []string
		for scanner.Scan() {
			line := scanner.Text()
			q = strings.Split(line, " ")
			plussing := 0
			for c, d := range q {
				if string(d[0]) == "(" {
					plussing++
					plussings = append(plussings, c)
				} else if string(d[len(d)-1]) == ")" {
					plussing--
				}

				if d == "+" {

					q[c-1] = "(" + q[c-1]
					plussing++
					plussings = append(plussings, c)
					if plussing > 0 && len(q)-2 == c {
						q[c+1] = q[c+1] + ")"
						plussing--
					} else {

					}
				} else if d == "*" && plussing > 0 {
					q[c-1] = q[c-1] + ")"
					plussing--

				}
			}

			z := ""
			for i, o := range q {
				if i == 0 {
					z = o
				} else {
					z = z + " " + o
				}

			}
			fmt.Println(z)

		}
	*/
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
	fmt.Println(tot)

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
