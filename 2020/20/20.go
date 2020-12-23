package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tile struct {
	label          string
	orientation1   [10][10]string
	orientation2   [10][10]string
	orientation3   [10][10]string
	orientation4   [10][10]string
	rotation       string
	unmatchedEdges int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	parseTileLabel := true
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
			tile = Tile{label: tileName, rotation: "N"}
		} else if line == "" {
			parseTileLabel = true
			tileMap[tileName] = tile
		} else {
			for l, z := range line {
				tile.orientation1[x][l] = string(z)

			}
			x++
		}

	}

	for z, r := range tileMap {
		or2 := rotate(r.orientation1)
		r.orientation2 = or2
		or3 := rotate(or2)
		r.orientation3 = or3
		or4 := rotate(or3)
		r.orientation4 = or4
		tileMap[z] = r
	}
	inner := 0
	for s := range tileMap {
		myBorders := getBorders(tileMap[s].orientation1)
		//fmt.Println(myBorders)
		totalTileMatches := 0
		for k, r := range tileMap {
			if s != k {
				found := false
				for _, m := range getBorders(r.orientation1) {
					cont, _ := ContainsString(myBorders, m)
					if cont {
						found = true
					}
				}
				for _, m := range getBorders(r.orientation2) {
					cont, _ := ContainsString(myBorders, m)
					if cont {
						found = true
					}
				}
				for _, m := range getBorders(r.orientation3) {
					cont, _ := ContainsString(myBorders, m)
					if cont {
						found = true
					}
				}
				for _, m := range getBorders(r.orientation4) {
					cont, _ := ContainsString(myBorders, m)
					if cont {
						found = true
					}
				}
				if found {
					totalTileMatches++
				}
			}
		}
		if totalTileMatches == 2 {
			fmt.Println("Found corners tiles for", s)
		} else if totalTileMatches == 3 {
			fmt.Println("Found edges", s)
		} else {
			fmt.Println("inner", s, totalTileMatches)
			inner++
		}
	}
	fmt.Println("inners", inner)

}

func rotate(matrix [10][10]string) (res [10][10]string) {
	n := len(matrix)
	x := n / 2
	y := n - 1
	for i := 0; i < x; i++ {
		for j := i; j < y-i; j++ {
			k := matrix[i][j]
			matrix[i][j] = matrix[y-j][i]
			matrix[y-j][i] = matrix[y-i][y-j]
			matrix[y-i][y-j] = matrix[j][y-i]
			matrix[j][y-i] = k
		}
	}
	res = matrix
	return
}

func getBorders(matri [10][10]string) (borders []string) {
	border1 := ""
	border2 := ""
	border3 := ""
	border4 := ""
	for i := 0; i < 10; i++ {
		border1 = border1 + matri[0][i]
		border2 = border2 + matri[i][0]
		border3 = border3 + matri[9][i]
		border4 = border4 + matri[i][9]
	}
	borders = append(borders, border1)
	borders = append(borders, border2)
	borders = append(borders, border3)
	borders = append(borders, border4)
	return
}
