package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Buddies struct {
	orbits     string
	indirectly []string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	orbitals := make(map[string]Buddies)

	for scanner.Scan() {
		line := scanner.Text()
		planets := strings.Split(line, ")")
		if _, ok := orbitals[planets[0]]; ok {
			//do something here
			//do nada
		} else {
			a := "NOTHING"
			var b []string
			setup := Buddies{orbits: a, indirectly: b}
			orbitals[planets[0]] = setup
		}

		if _, ok := orbitals[planets[1]]; ok {
			//do something here
			//do nada
			tem := orbitals[planets[1]]
			tem.orbits = planets[0]
			orbitals[planets[1]] = tem
		} else {
			var b []string
			setup := Buddies{orbits: planets[0], indirectly: b}
			orbitals[planets[1]] = setup
		}

	}
	orbits := 0
	indirect := 0
	for a, val := range orbitals {
		fmt.Println("Plantes ", a, " orbits ", val.orbits)
		if val.orbits != "NOTHING" {
			orbits = orbits + 1
			indirect = indirect + followTheOrbits(orbitals, val.orbits)
		}
	}
	fmt.Println(orbits, indirect)
	fmt.Println("total", orbits+indirect)
	me := countTheOrbits(orbitals, "YOU")
	santa := countTheOrbits(orbitals, "SAN")
	fmt.Println(me)
	fmt.Println(santa)
	for i, a := range me {
		cont, index := contains(santa, a)
		if cont {
			fmt.Println("SantaRange,", index, "Merange", i)
			fmt.Println(index + i - 2)
			fmt.Println(a)

			break

		}

	}

}

func followTheOrbits(orbs map[string]Buddies, key string) (result int) {

	if orbs[key].orbits == "NOTHING" {
		return 0
	}
	return followTheOrbits(orbs, orbs[key].orbits) + 1

}

func countTheOrbits(orbs map[string]Buddies, key string) (result []string) {

	if orbs[key].orbits == "NOTHING" {
		return result
	}
	result = append(result, key)
	return append(result, countTheOrbits(orbs, orbs[key].orbits)...)

}

func contains(s []string, e string) (bool, int) {
	for i, a := range s {
		if a == e {
			return true, i
		}
	}
	return false, 0
}
