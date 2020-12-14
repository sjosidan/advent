package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Memory struct {
	store   int64
	mString string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	memoryMap := make(map[int64]*Memory)
	memoryMap2 := make(map[string]*Memory)

	bitMask := ""
	var newProgram bool
	fmt.Println(newProgram)
	for scanner.Scan() {

		line := scanner.Text()
		if line[:4] == "mask" {
			newProgram = true
			bitMask = line[7:]
		} else {
			newProgram = false
			splitter := strings.Split(line, " ")

			inputValue := getBinaryString(splitter[len(splitter)-1])
			memoryAdress := getBinaryString(splitter[0][4 : len(splitter[0])-1])
			memAdr, _ := strconv.ParseInt(memoryAdress, 2, 64)

			memValue := maskMe(inputValue, bitMask)
			numericValue, _ := strconv.ParseInt(memValue, 2, 64)

			maskedAdresses := maskAdress(memoryAdress, bitMask)
			val2, _ := strconv.Atoi(splitter[len(splitter)-1])

			//GET ALL ADRESSES
			writeSlice := perms(maskedAdresses)

			if _, ok := memoryMap[memAdr]; ok {
				memoryMap[memAdr].mString = memValue
				memoryMap[memAdr].store = numericValue

			} else {

				m := Memory{mString: memValue, store: numericValue}
				memoryMap[memAdr] = &m

			}
			for _, mAdr := range writeSlice {
				if _, ok := memoryMap2[mAdr]; ok {
					memoryMap2[mAdr].mString = inputValue
					memoryMap2[mAdr].store = int64(val2)
				} else {
					m2 := Memory{mString: inputValue, store: int64(val2)}
					memoryMap2[mAdr] = &m2

				}
			}
		}
	}
	sum := int64(0)

	for _, v := range memoryMap {
		sum = sum + v.store
	}

	sum2 := float64(0)

	for k, v := range memoryMap2 {
		floates := strings.Count(k, "X")
		sum2 = sum2 + (float64(v.store) * math.Pow(2, float64(floates)))
	}

	fmt.Println("Answer 1 Sum", sum)
	fmt.Println("Answer 2 Sum", int64(sum2))
}

func getBinaryString(numericString string) (result string) {

	intReps, _ := strconv.Atoi(numericString)
	a := strconv.FormatInt(int64(intReps), 2)
	result = lpad(a, "0", 36)
	return
}

func lpad(s string, pad string, plength int) string {
	for i := len(s); i < plength; i++ {
		s = pad + s
	}
	return s
}

func maskMe(s string, mask string) (result string) {
	for i := 0; i < len(mask); i++ {
		switch string(mask[i]) {
		case "X":
			result = result + string(s[i])
		default:
			result = result + string(mask[i])

		}
	}

	return
}

func perms(adr string) (result []string) {
	xses := strings.Count(adr, "X")
	var bitArray []string
	comb := math.Pow(2, float64(xses))
	for i := 0; i < int(comb); i++ {
		bitArray = append(bitArray, lpad(strconv.FormatInt(int64(i), 2), "0", xses))
	}
	for _, bitMap := range bitArray {
		xCount := 0
		adrCopy := adr
		for i, z := range adr {

			if string(z) == "X" {
				//REPLACE WITH FIRST VALUE FROM bitMap
				adrCopy = replaceAtIndex(adrCopy, rune(bitMap[xCount]), i)
				xCount++
			}
		}
		result = append(result, adrCopy)
	}
	return
}

func maskAdress(s string, mask string) (result string) {

	for i := 0; i < len(mask); i++ {
		switch string(mask[i]) {
		case "0":
			result = result + string(s[i])
		default:
			result = result + string(mask[i])

		}
	}
	return
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
