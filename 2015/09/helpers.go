package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
)

func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsString(s []string, e string) (bool, int) {
	for i, a := range s {
		if a == e {
			return true, i
		}
	}
	return false, 0
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

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func Abss(a int, b int) (result float64) {
	result = math.Abs(float64(a)) + math.Abs(float64(b))
	return
}

func PermutationsInt(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func PermutationsString(testStr string) []string {
	var n func(testStr []rune, p []string) []string
	n = func(testStr []rune, p []string) []string {
		if len(testStr) == 0 {
			return p
		} else {
			result := []string{}
			for _, e := range p {
				result = append(result, join([]rune(e), testStr[0])...)
			}
			return n(testStr[1:], result)
		}
	}

	output := []rune(testStr)
	return n(output[1:], []string{string(output[0])})
}

func join(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}

func ConvertStringSliceToIntSlice(strings []string) []int {
	ints := make([]int, len(strings))

	for i, s := range strings {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetSubstringsOfLength(bigString string, length int) (substrings []string) {

	for i, _ := range bigString {
		if i+length+1 < len(bigString)+1 {
			substrings = append(substrings, bigString[i:i+length])
		}
	}

	return

}

/*
func HeapPermutation(a []string, size int) (result []string) {
	if size == 1 {
		result = a
		fmt.Println(a)
	}

	for i := 0; i < size; i++ {
		HeapPermutation(a, size-1)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
	return
}

*/
var ans [][]string

func PermuteStringsSlice(nums []string) [][]string {
	ans = make([][]string, 0)
	cp := make([]string, len(nums))
	copy(cp, nums)
	ans = append(ans, cp)

	if len(nums) < 2 {
		return ans
	}

	heap(nums, len(nums), len(nums))
	return ans
}

func heap(nums []string, n, l int) {

	for i := 0; ; i++ {
		if n > 2 {
			heap(nums, n-1, l)
		}
		if i == n-1 {
			break
		}
		if n%2 == 0 {
			swap(nums, i, n-1)
		} else {
			swap(nums, 0, n-1)
		}
		cp := make([]string, len(nums))
		copy(cp, nums)
		ans = append(ans, cp)
	}
}

func swap(nums []string, i, j int) {
	tmp := nums[j]
	nums[j] = nums[i]
	nums[i] = tmp
}
