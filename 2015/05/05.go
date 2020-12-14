package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	niceKids1 := 0
	niceKids2 := 0

	naughtKids1 := 0
	naughtKids2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		niceKid := amINice(line)
		if niceKid {
			niceKids1++
		} else {
			naughtKids1++
		}

		niceKid = amINice2(line)
		if niceKid {
			niceKids2++
		} else {
			naughtKids2++
		}

	}
	fmt.Println("Nice kids first Rules", niceKids1)
	fmt.Println("Nice kids second Rules", niceKids2)
}

func amINice(child string) (nice bool) {
	vowelCounter := 0
	doubleCheck := false
	lastRune := ""
	illegals := false
	for _, l := range child {
		doubleDig := lastRune + string(l)
		if string(l) == "a" ||
			string(l) == "e" ||
			string(l) == "i" ||
			string(l) == "o" ||
			string(l) == "u" {
			vowelCounter++
		}

		if lastRune == string(l) {
			doubleCheck = true
		}

		if doubleDig == "ab" ||
			doubleDig == "cd" ||
			doubleDig == "pq" ||
			doubleDig == "xy" {
			illegals = true
		}
		lastRune = string(l)

	}
	if vowelCounter > 2 && !illegals && doubleCheck {
		nice = true
	}
	return
}

func amINice2(child string) (nice bool) {
	substringsTwo := GetSubstringsOfLength(child, 2)
	substringsThree := GetSubstringsOfLength(child, 3)
	repeatingChar := false
	multiDouble := false
	nice = false
	for i, ss := range substringsTwo {
		containsCount := strings.Count(child, ss)
		if 1 < containsCount {
			if containsCount == 2 && len(substringsTwo) > i+1 && i != 0 {
				if ss != substringsTwo[i-1] && ss != substringsTwo[i+1] {
					multiDouble = true
				}
			}
		}
	}

	for _, ss := range substringsThree {
		if ss[0] == ss[2] {
			repeatingChar = true
		}
	}
	if repeatingChar && multiDouble {
		nice = true
	}

	return
}
