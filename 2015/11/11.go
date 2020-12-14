package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//line := scanner.Text()
	}

	//numArray := []int{0, 0, 0, 0, 0, 0, 0, 0}
	//vzbxkghb
	numArray := []int{21, 25, 1, 23, 10, 6, 7, 1}
	//abcdefgh
	//numArray := []int{0, 1, 2, 3, 4, 5, 6, 7}

	found := false
	var secondPass []int
	for {
		dup := checkValidDuplicates(numArray)
		follow := checkValidFollow(numArray)
		noIllegals := checkforIllegas(numArray)
		if dup && follow && noIllegals {
			if !found {
				fmt.Println("First Pass:")
				for _, c := range numArray {
					fmt.Print(translateToChar(c))
				}
				fmt.Println("")
				found = true
			} else {
				secondPass = numArray
				break
			}
		}
		numArray = increaseArray(numArray)

	}

	fmt.Println("Second Pass")
	for _, c := range secondPass {
		fmt.Print(translateToChar(c))
	}
	fmt.Println("")
}

func checkValidDuplicates(numArray []int) (valid bool) {
	valids := 0
	for i := 1; i < len(numArray); i++ {
		if numArray[i] == numArray[i-1] {
			valids++
			i++
		}
	}
	if valids > 1 {
		valid = true
	}
	return
}

func checkValidFollow(numArray []int) (valid bool) {
	for i := 2; i < len(numArray); i++ {
		if numArray[i]-numArray[i-1] == 1 && numArray[i-1]-numArray[i-2] == 1 {
			valid = true
			break
		}
	}
	return
}

func checkforIllegas(numArray []int) (valid bool) {
	valid = true
	for i := 0; i < len(numArray); i++ {
		//i, o, l
		//
		if numArray[i] == 8 || numArray[i] == 11 || numArray[i] == 14 {
			valid = false
			break
		}
	}

	return
}

func increaseArray(numArray []int) (res []int) {
	if numArray[len(numArray)-1] == 25 {
		numArray[len(numArray)-1] = 0
		if numArray[len(numArray)-2] == 25 {
			numArray[len(numArray)-2] = 0
			if numArray[len(numArray)-3] == 25 {
				numArray[len(numArray)-3] = 0
				if numArray[len(numArray)-4] == 25 {
					numArray[len(numArray)-4] = 0
					if numArray[len(numArray)-5] == 25 {
						numArray[len(numArray)-5] = 0
						if numArray[len(numArray)-6] == 25 {
							numArray[len(numArray)-6] = 0
							if numArray[len(numArray)-7] == 25 {
								numArray[len(numArray)-7] = 0
							} else {
								numArray[len(numArray)-7]++
							}
						} else {
							numArray[len(numArray)-6]++
						}
					} else {
						numArray[len(numArray)-5]++
					}
				} else {
					numArray[len(numArray)-4]++
				}
			} else {
				numArray[len(numArray)-3]++
			}
		} else {
			numArray[len(numArray)-2]++
		}
	} else {
		numArray[len(numArray)-1]++
	}
	res = numArray
	return numArray
}

func translateToChar(i int) string {
	switch i {
	case 0:
		return "a"
	case 1:
		return "b"
	case 2:
		return "c"
	case 3:
		return "d"
	case 4:
		return "e"
	case 5:
		return "f"
	case 6:
		return "g"
	case 7:
		return "h"
	case 8:
		return "i"
	case 9:
		return "j"
	case 10:
		return "k"
	case 11:
		return "l"
	case 12:
		return "m"
	case 13:
		return "n"
	case 14:
		return "o"
	case 15:
		return "p"
	case 16:
		return "q"
	case 17:
		return "r"
	case 18:
		return "s"
	case 19:
		return "t"
	case 20:
		return "u"
	case 21:
		return "v"
	case 22:
		return "w"
	case 23:
		return "x"
	case 24:
		return "y"
	case 25:
		return "z"
	default:
	}
	return ""
}
