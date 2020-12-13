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
				if !Contains(steps, k.departs) {
					steps = append(steps, k.departs)
					step = Sum(steps)
				}
			} else {
				works = false
			}
		}
		if works {
			break
		}
		i = i + step
	}
	fmt.Print("Time\t")
	for _, k := range busses {
		fmt.Print("Bus ", k.name, "\t")
	}
	fmt.Println("")

	for j := i; j < i+busses[len(busses)-1].globalOffset+1; j++ {
		fmt.Print(j, "\t")
		for _, k := range busses {
			if j%k.departs == 0 {
				fmt.Print("D\t")
			} else {
				fmt.Print(".\t")
			}
		}
		fmt.Println("")
	}
	fmt.Println("Answer B:", i)
}

func taskA(now int, busses []Bus) {
	wait := now
	foundbus := false
	for {
		for k, v := range busses {
			if wait%v.departs == 0 {
				fmt.Println("Answer A: Can depart bus", k, "at", wait, "ans", (wait-now)*v.departs)
				foundbus = true
			}
		}
		if foundbus {
			break
		}
		wait++
	}
}
