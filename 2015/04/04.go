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
	seed := ""
	for scanner.Scan() {
		seed = scanner.Text()
	}
	hashString := ""
	ans := 0
	for {
		fully := seed + strconv.Itoa(ans)
		hashString = GetMD5Hash(fully)

		if strings.HasPrefix(hashString, "000000") {
			break
		}
		ans++
	}
	fmt.Println(hashString)
	fmt.Println("Answer", ans)
}
