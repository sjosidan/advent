package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sumLen := 0
	sumRemoved := 0
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Printf(line)
		updated := line

		countEscQuotes := strings.Count(line, "\\\"")
		updated = strings.Replace(updated, "\\\"", "-", -1)
		//fmt.Println(updated)

		countQuotes := strings.Count(updated, "\"")
		updated = strings.Replace(updated, "\"", "", -1)
		//fmt.Println(updated)

		countBackSlash := strings.Count(updated, "\\\\")
		updated = strings.Replace(updated, "\\\\", "-", -1)
		//fmt.Println(updated)

		countHex := strings.Count(updated, "\\x")
		updated = strings.Replace(updated, "\\x", "-", -1)
		//fmt.Println(updated)

		a := len(line) - (countHex * 3) - countEscQuotes - countBackSlash - countQuotes

		fmt.Println(updated, a, len(line))
		sumLen = sumLen + len(line)
		sumRemoved = sumRemoved + a
	}
	fmt.Println(sumLen)
	fmt.Println(sumRemoved)
	fmt.Println(sumLen - sumRemoved)
	fmt.Println("nxzo\"hf\\xp")

}
