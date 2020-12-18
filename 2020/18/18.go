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
			z, _ := evalParent(k[a.start:a.end])
			reps[k[a.start:a.end]] = z
			keys = append(keys, key{key: k[a.start:a.end], val: z})
		}
		for _, i := range keys {
			for _, r := range keys {
				if i.key != r.key {
					if strings.Contains(r.key, i.key) {
						nn, _ := evalParent(strings.Replace(r.key, i.key, strconv.Itoa(i.val), -1))
						reps[r.key] = nn
						delete(reps, i.key)

					}
				}
			}
		}
		a := line
		for i, k := range reps {
			a = strings.Replace(a, i, strconv.Itoa(k), -1)
		}
		loc, _ := evalParent(a)
		tot = tot + loc
		fmt.Println(a, loc)
	}
	fmt.Println(tot)

}

func evalParent(line string) (tot int, leng int) {
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
			n, _ := strconv.Atoi(unescp[i+1])
			tot = tot + n
		} else {
			n, _ := strconv.Atoi(unescp[i+1])
			tot = tot * n
		}
		i++
		leng = i

	}
	return
}
