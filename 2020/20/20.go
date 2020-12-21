package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tile struct {
	label          string
	orientation1   [10][10]string
	border1        string
	border2        string
	border3        string
	border4        string
	rotation       string
	unmatchedEdges int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	borderMap := make(map[string][]Tile)
	ugg := make(map[string]int)

	parseTileLabel := true
	//newTile := false
	tileName := ""
	var tiles []string
	x := 0
	tileMap := make(map[string]Tile)
	var tile Tile
	for scanner.Scan() {
		line := scanner.Text()

		if parseTileLabel {
			tileName = line
			parseTileLabel = false
			tiles = append(tiles, tileName)
			x = 0
			tile = Tile{label: tileName}

		} else if line == "" {
			parseTileLabel = true
			tileMap[tileName] = tile

		} else {
			tile.border2 = tile.border2 + string(line[0])
			tile.border3 = tile.border3 + string(line[9])

			if x == 0 {
				tile.border1 = line

			}
			if x == 9 {
				tile.border4 = Reverse(line)
				tile.rotation = "N"
				tile.border2 = Reverse(tile.border2)
				borderMap[tile.border1] = append(borderMap[tile.border1], tile)
				borderMap[tile.border2] = append(borderMap[tile.border2], tile)
				borderMap[tile.border3] = append(borderMap[tile.border3], tile)
				borderMap[tile.border4] = append(borderMap[tile.border4], tile)

			}

			for l, z := range line {
				tile.orientation1[x][l] = string(z)

			}
			x++
		}

	}

	for z, r := range tileMap {
		fmt.Println(z, r.rotation)
		fmt.Println(r.border1)
		fmt.Println(r.border2)
		fmt.Println(r.border3)
		fmt.Println(r.border4)
		for _, rr := range r.orientation1 {
			fmt.Println(rr)
		}
	}

	for a, r := range borderMap {
		if len(r) == 1 {
			for _, z := range r {
				fmt.Println(a, z.label)
				ugg[z.label]++

			}
		}
	}
	fmt.Println(ugg)

	for z, r := range ugg {
		if r == 2 {
			fmt.Println(z, "must be corner")
		}
		if r == 1 {
			fmt.Println(z, "must be edge")
		}
	}

	//fmt.Println(borderMap)
	//	perms := PermuteStringsSlice(tiles)
	//fmt.Println(perms)
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
