package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Connection struct {
	cityA    string
	cityB    string
	distance int
}

//Dublin to Belfast = 141
func main() {
	cities := make(map[string][]Connection)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		conn := strings.Split(line, " ")
		dist, _ := strconv.Atoi(conn[4])
		connection := Connection{cityA: conn[0], cityB: conn[2], distance: dist}

		cities[conn[0]] = append(cities[conn[0]], connection)

		cities[conn[2]] = append(cities[conn[2]], connection)

	}

	citySlice := make([]string, len(cities))
	i := 0
	for k := range cities {
		citySlice[i] = k
		i++
	}

	routes := PermuteStringsSlice(citySlice)
	shortestRoute := 9999999
	longestRoute := 0
	var minRoute []string
	var maxRoute []string

	for _, route := range routes {
		length := 0
		for i, city := range route {
			if len(route)-1 != i {
				for _, con := range cities[city] {
					if con.cityA == route[i+1] || con.cityB == route[i+1] {
						length = length + con.distance
						break
					}
				}
			}
		}
		if length < shortestRoute {
			shortestRoute = length
			minRoute = route
		}
		if length > longestRoute {
			longestRoute = length
			maxRoute = route
		}

	}
	fmt.Println("Shortest", minRoute, shortestRoute)
	fmt.Println("Longest", maxRoute, longestRoute)
}
