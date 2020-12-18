package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		terms := strings.Split(line, " ")
		//leftCounter := 0
		//startIndex := 0
		//var evalMe []string
		ll := len(terms)
		fmt.Println(line)
		var newTerms []string
		for i := 0; i < ll; i++ {
			if strings.HasPrefix(terms[i], "(") {

				newSlice := terms[i:]
				_, res := parseParant(newSlice, i)
				newTerms = append(newTerms, strconv.Itoa(res))
			} else {
				newTerms = append(newTerms, terms[i])
			}

		}
		//fmt.Println(evalParent(evalMe))
	}
}

func parseParant(newSlice []string, startIndex int) (res int, ll int) {
	fmt.Println("solve,", newSlice)
	begIndex := startIndex
	endIndex := startIndex
	var mySlice []string
	for j := 0; j < len(newSlice); j++ {
		if strings.HasPrefix(newSlice[j], "(") && j != 0 {
			kkslice := newSlice[:j]
			_, z := parseParant(kkslice, j)
			j = j + z

		}
		if strings.HasSuffix(newSlice[j], ")") {
			fmt.Println(j, newSlice)
			val := evalParent(newSlice[:j+1])
			j = j + len(newSlice[:j+1])
			return val, ll
		}
		mySlice = append(mySlice, newSlice[j])

	}

	fmt.Println(begIndex, endIndex)

	return 0, endIndex - begIndex

}

func evalParent(line []string) (tot int) {
	fmt.Println("-----", line)
	tot, _ = strconv.Atoi(strings.Replace(line[0], "(", "", -1))
	var unescp []string
	for _, z := range line {
		n := strings.Replace(z, ")", "", -1)
		n = strings.Replace(n, "(", "", -1)
		unescp = append(unescp, n)
	}
	fmt.Println(unescp)
	for i := 1; i < len(unescp)-1; i++ {

		if line[i] == "+" {
			n, _ := strconv.Atoi(unescp[i+1])
			tot = tot + n
		} else {
			n, _ := strconv.Atoi(unescp[i+1])
			tot = tot * n
		}
		i++
	}

	return
}

/*
func evalBlock(line []string) (tot int, blockSize int) {
	fmt.Println(line)
	var expr []string

	for i := 0; i < len(line)-1; i++ {
		fmt.Println(line[i], i)
		if strings.HasSuffix(line[i], ")") {
			fmt.Println("close")
			expr = append(expr, line[i])
			expr[0] = strings.Replace(line[0], "(", "", -1)
			expr[len(expr)-1] = strings.Replace(line[i], ")", "", -1)
			fmt.Println(expr, line[i])
			tot = evalParent(expr)
			return tot, len(expr)

		}
		if strings.HasPrefix(line[i], "(") && i != 0 {
			fmt.Println("open")
			n, inc := evalBlock(line[i:])
			expr = append(expr, strconv.Itoa(n))
			fmt.Println("---", n)
			i = i + inc
			tot = tot + i
		}

		expr = append(expr, line[i])
	}
	return
}
*/
