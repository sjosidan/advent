package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	/*
	*
	* Trimmed input , removed all ) and commas for easier parsing.
	 */
	scanner := bufio.NewScanner(os.Stdin)
	allargenMap := make(map[string][]string)
	ingridientMap := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, " (contains ")
		allargens := strings.Split(s[1], " ")
		ingredients := strings.Split(s[0], " ")
		for _, a := range allargens {
			if val, ok := allargenMap[a]; ok {
				var newSlice []string
				for _, exists := range val {
					for _, ing := range ingredients {
						if exists == ing {
							newSlice = append(newSlice, ing)
						}
					}
				}
				allargenMap[a] = newSlice
			} else {
				allargenMap[a] = ingredients
			}
		}
		for _, ing := range ingredients {
			if ing != "" {
				ingridientMap[ing]++
			}
		}
	}
	//fmt.Println(allargenMap)
	figureOutWhich(allargenMap)
	//fmt.Println(ingridientMap)
	getCount(ingridientMap, allargenMap)
}

func getCount(ingMap map[string]int, aMap map[string][]string) {
	tot := 0
	for a, count := range ingMap {
		found := false
		for _, s := range aMap {
			for _, contains := range s {
				if contains == a {
					found = true
				}
			}
		}

		if !found {
			tot = tot + count
		}
	}
	fmt.Println("")
	fmt.Println("totals,", tot)

}

func figureOutWhich(a map[string][]string) {
	finalMap := make(map[string]string)
	z := 0
	remove := ""
	for {

		for typ, allargen := range a {
			if len(allargen) == 1 {
				finalMap[typ] = allargen[0]
				//REMOVE
				remove = allargen[0]
				break
			}
		}

		//PURGE FROM MAP
		for key, z := range a {

			cont, i := ContainsString(z, remove)
			if cont {
				a[key] = RemoveIndex(z, i)
			}

		}
		z++

		if z == 100 {
			break
		}
	}
	var keys []string

	for k, _ := range finalMap {
		keys = append(keys, k)
	}
	sort.Sort(Alphabetic(keys))
	for _, k := range keys {
		fmt.Print(finalMap[k] + ",")
	}

}
