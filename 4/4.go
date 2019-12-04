package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := 356261
	b := 846303
	curr := a
	count := 0
	for curr < b {
		//	fmt.Println(curr)
		curr++
		si := strconv.Itoa(curr)
		legit := true
		v1 := strings.Count(si, "44") + strings.Count(si, "444")
		v2 := strings.Count(si, "55") + strings.Count(si, "555")
		v3 := strings.Count(si, "66") + strings.Count(si, "666")
		v4 := strings.Count(si, "77") + strings.Count(si, "777")
		v5 := strings.Count(si, "88") + strings.Count(si, "888")
		v6 := strings.Count(si, "99") + strings.Count(si, "999")
		if v1 == 1 || v2 == 1 || v3 == 1 || v4 == 1 || v5 == 1 || v6 == 1 {
			lastNum := int(si[0])
			for _, i := range si {
				a = int(i)
				if a < lastNum {
					legit = false
				}
				lastNum = a
			}
		} else {
			legit = false
		}
		if legit {
			count = count + 1
			fmt.Println(si)
			fmt.Println(v1, v2, v3, v4, v5, v6)
		}
	}
	fmt.Println(count)
}
