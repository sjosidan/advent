package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lamp struct {
	x          int
	y          int
	brightness int
	on         bool
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lampGrid := make(map[string]*Lamp)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			key := strconv.Itoa(i) + "-" + strconv.Itoa(j)
			lamp := Lamp{x: i, y: j, on: false}
			lampGrid[key] = &lamp
		}
	}

	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	for scanner.Scan() {
		line := scanner.Text()

		args := strings.Split(line, " ")
		switch len(args) {
		case 4:
			mins := strings.Split(args[1], ",")
			minX, _ = strconv.Atoi(mins[0])
			minY, _ = strconv.Atoi(mins[1])

			maxs := strings.Split(args[3], ",")
			maxX, _ = strconv.Atoi(maxs[0])
			maxY, _ = strconv.Atoi(maxs[1])
			//	fmt.Println("modified", minX, minY, maxX, maxY)
			for i := minX; i <= maxX; i++ {
				for j := minY; j <= maxY; j++ {
					key := strconv.Itoa(i) + "-" + strconv.Itoa(j)
					lampGrid[key].on = !lampGrid[key].on
					lampGrid[key].brightness = lampGrid[key].brightness + 2

				}
			}
		case 5:
			//fmt.Println("turn on off")
			mins := strings.Split(args[2], ",")
			minX, _ = strconv.Atoi(mins[0])
			minY, _ = strconv.Atoi(mins[1])

			maxs := strings.Split(args[4], ",")
			maxX, _ = strconv.Atoi(maxs[0])
			maxY, _ = strconv.Atoi(maxs[1])
			for i := minX; i <= maxX; i++ {
				for j := minY; j <= maxY; j++ {
					key := strconv.Itoa(i) + "-" + strconv.Itoa(j)
					if args[1] == "on" {
						lampGrid[key].on = true
						lampGrid[key].brightness = lampGrid[key].brightness + 1

					} else {
						lampGrid[key].on = false
						if lampGrid[key].brightness > 0 {
							lampGrid[key].brightness = lampGrid[key].brightness - 1
						}

					}
				}
			}
		default:
		}
	}
	lampsOn := 0
	brightNess := 0
	for _, v := range lampGrid {
		if v.on {
			lampsOn++
		}
		if v.brightness > 0 {
			brightNess = brightNess + v.brightness
		}
	}
	fmt.Println("Lamps lit", lampsOn)
	fmt.Println("Brightness", brightNess)

}
