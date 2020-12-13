package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bus struct {
	name         string
	departs      int
	globalOffset int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	firstLine := true
	var busses []Bus
	var now int
	for scanner.Scan() {

		line := scanner.Text()
		if firstLine {
			now, _ = strconv.Atoi(line)
			firstLine = false
		} else {
			lines := strings.Split(line, ",")
			offset := 0
			for _, l := range lines {
				if l != "x" {
					depart, _ := strconv.Atoi(l)
					bus := Bus{name: l, departs: depart, globalOffset: offset}
					busses = append(busses, bus)

				}
				offset++
			}

		}
	}

	fmt.Println(now)
	fmt.Println(busses)
	taskA(now, busses)
	taskB(busses)

}
func taskB(busses []Bus) {
	i := 0
	var steps []int
	steps = append(steps, busses[0].departs)
	step := busses[0].departs

	for {
		works := true
		for _, k := range busses {
			if works && ((i+k.globalOffset)%(k.departs) == 0) {
				if !contains(steps, k.departs) {
					steps = append(steps, k.departs)
					fmt.Println(i, " Works for Bus ", k.name, " ")
					step = sum(steps)
				}
			} else {
				works = false
			}
		}
		if works {
			fmt.Println("Fine at", i)
			break
		}
		i = i + step
	}
}

func taskA(now int, busses []Bus) {
	wait := now
	foundbus := false
	for {
		for k, v := range busses {
			if wait%v.departs == 0 {
				fmt.Println("Can depart bus", k, " at ", wait, " ans ", (wait-now)*v.departs)
				foundbus = true
			}
		}
		if foundbus {
			break
		}
		wait++
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func sum(array []int) int {
	result := 1
	for _, v := range array {
		result *= v
	}
	return result
}
