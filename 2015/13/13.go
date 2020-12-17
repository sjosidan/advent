package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Relation struct {
	person string
	value  int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	opinions := make(map[string][]Relation)
	for scanner.Scan() {
		line := scanner.Text()
		rule := strings.Split(line, " ")
		unit := 1
		person1 := rule[0]
		if rule[2] == "lose" {
			unit = -1
		}

		person2 := rule[10][:len(rule[10])-1]
		amount, _ := strconv.Atoi(rule[3])

		if _, ok := opinions[person1]; ok {
			//do something here
			r := Relation{person: person2, value: unit * amount}
			opinions[person1] = append(opinions[person1], r)
		} else {
			r := Relation{person: person2, value: unit * amount}
			opinions[person1] = append(opinions[person1], r)
		}
	}
	var people []string

	for p, _ := range opinions {
		people = append(people, p)
	}
	options := PermuteStringsSlice(people)

	maxVal := 0
	var bestSeat []string
	var bestCirc []string
	for _, k := range options {

		circularTable := makeStringCirc(k)
		//fmt.Println(circularTable)
		seatVal := 0
		for j, l := range circularTable {

			for _, rel := range opinions[l] {
				if j == len(circularTable)-1 {
					if rel.person == circularTable[0] {
						seatVal = seatVal + rel.value
					}

				} else {
					if rel.person == circularTable[j+1] {
						seatVal = seatVal + rel.value
					}
				}
			}
		}

		if seatVal > maxVal {
			maxVal = seatVal
			bestSeat = k
			bestCirc = circularTable
		}
	}

	fmt.Println("BEST:", bestSeat, maxVal, bestCirc)
}

func makeStringCirc(listA []string) (listB []string) {

	s := append([]string{}, listA...)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	listB = append(listA, s[1:]...)
	listB = append([]string{s[0]}, listB...)
	return
}
