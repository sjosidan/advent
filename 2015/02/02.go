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
	paperNeeded := 0
	ribbonNeeded := 0
	for scanner.Scan() {
		line := scanner.Text()
		dimensions := strings.Split(line, "x")

		l, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		h, _ := strconv.Atoi(dimensions[2])

		intSlice := []int{l * w, w * h, h * l}
		min, _ := MinMax(intSlice)

		surfaceArea := (2 * l * w) + (2 * w * h) + (2 * h * l)
		volume := l * w * h
		paperNeeded = paperNeeded + surfaceArea + min

		//Ribbon

		//sideLength := ConvertStringSliceToIntSlice(dimensions)
		circumferences := []int{l + l + w + w, w + w + h + h, h + h + l + l}
		minCircum, _ := MinMax(circumferences)
		ribbonNeeded = ribbonNeeded + minCircum + volume
	}
	fmt.Println("Total Amount of paper needed ", paperNeeded)
	fmt.Println("Total Amount of ribbon needed ", ribbonNeeded)
}
