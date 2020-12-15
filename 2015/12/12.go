package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tot := 0
	for scanner.Scan() {
		line := scanner.Text()
		alles := strings.Split(line, ",")
		fmt.Println(len(alles))
		for _, a := range alles {

			z := strings.Split(a, ":")
			for _, r := range z {
				r1 := strings.Replace(r, "[", "", -1)
				r2 := strings.Replace(r1, "]", "", -1)

				r3 := strings.Replace(r2, "}", "", -1)
				r4 := strings.Replace(r3, "{", "", -1)
				fmt.Println(r4)
				val, err := strconv.Atoi(r4)
				if err == nil {
					fmt.Println(val)
					tot = tot + val
				}

			}
		}
		//jsonData := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

		var v interface{}
		json.Unmarshal([]byte(line), &v)
		data := v.(map[string]json.RawMessage)
		fmt.Println((data))

	}
	fmt.Println(tot)
}
