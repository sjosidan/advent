package main

import "fmt"

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func Sum(array []int) int {
	result := 1
	for _, v := range array {
		result *= v
	}
	return result
}

func HW() {
	fmt.Println("hello world")
}
