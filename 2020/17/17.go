package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Cube struct {
	x      int
	y      int
	z      int
	w      int
	active bool
	state  string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cubeMap := make(map[string]Cube)

	zl := 0
	yl := 0
	wl := 0
	zs := []int{0}
	zSlice := make([][]int, 3)
	for scanner.Scan() {

		line := scanner.Text()
		for i, a := range line {
			active := false
			on := 0
			if string(a) == "#" {
				active = true
				on = 1
			}
			c := Cube{x: i, y: yl, z: zl, w: wl, active: active, state: string(a)}
			key := strconv.Itoa(i) + "-" + strconv.Itoa(yl) + "-" + strconv.Itoa(zl) + "-" + strconv.Itoa(wl)
			cubeMap[key] = c
			zSlice[zl] = append(zSlice[zl], on)
		}
		yl++

	}
	for k, v := range cubeMap {
		fmt.Println(k, v)
	}
	fmt.Println("")
	fmt.Println("")

	for _, z := range zs {
		fmt.Println("z = ", z)
		for r, v := range zSlice[0] {
			if r%3 == 0 {
				fmt.Println()
			}
			fmt.Print(v)
		}
	}
	fmt.Println("")

	min, max := MinMax(zs)
	fmt.Println("new zs", min-1, max+1)

	initZ := -1
	initX := 0
	initW := -1
	//amActive := false
	k := 0
	for {
		x1, x2, y1, y2, z1, z2, w1, w2 := getLowestEncirclingCube(cubeMap)
		fmt.Println(getLowestEncirclingCube(cubeMap))
		initZ = z1
		cubeMap2 := make(map[string]Cube)
		for {
			initW = w1

			for {
				initX = x1

				for {
					initY := y1

					for {
						meKey := strconv.Itoa(initX) + "-" + strconv.Itoa(initY) + "-" + strconv.Itoa(initZ) + "-" + strconv.Itoa(initW)
						activeNeig := 0
						amIactive := false

						key := strconv.Itoa(initX) + "-" + strconv.Itoa(initY) + "-" + strconv.Itoa(initZ) + "-" + strconv.Itoa(initW)

						if val, ok := cubeMap[meKey]; ok {
							if val.active {
								amIactive = true
							}
						}
						a := getNeighbours(initX, initY, initZ, initW)
						for _, n := range a {
							if val, ok := cubeMap[n]; ok {
								//do something here
								if val.active {
									activeNeig++
								}
							}
						}
						if amIactive {
							if activeNeig == 3 || activeNeig == 2 {
								//fmt.Println(initX, initY, initZ, "keep Active")
								amIactive = true
								c := Cube{x: initX, y: initY, z: initY, active: amIactive, state: ""}
								cubeMap2[key] = c

							}
						} else {
							if activeNeig == 3 {
								//fmt.Println(initX, initY, initZ, "Turn active")
								amIactive = true
								c := Cube{x: initX, y: initY, z: initY, active: amIactive, state: ""}
								cubeMap2[key] = c

							}
						}
						//fmt.Println(initX, initY, initZ, initW)

						if initY == y2 {
							break
						}
						initY++
					}
					if initX == x2 {
						break
					}
					initX++

				}
				if initW == w2 {
					break
				}
				initW++
			}
			if initZ == z2 {
				cubeMap = cubeMap2

				break

			}
			initZ++

		}
		if k == 6 {
			break
		}
		k++
		totLights := 0
		for _, v := range cubeMap2 {
			if v.active {
				totLights++
			}
		}
		fmt.Println(k, "lit", totLights)
	}

	/*
		for i, a := range line {
				active := false
				on := 0
				if string(a) == "#" {
					active = true
					on = 1
				}
				c := Cube{x: i, y: yl, z: zl, active: active, state: string(a)}
				key := strconv.Itoa(i) + "-" + strconv.Itoa(yl) + "-" + strconv.Itoa(zl)
				cubeMap[key] = c
				zSlice[zl] = append(zSlice[zl], on)
			}
			yl++
	*/

}

func getNeighbours(x int, y int, z int, w int) (neighbours []string) {
	for i := x - 1; i < x+2; i++ {
		for j := y - 1; j < y+2; j++ {

			for k := z - 1; k < z+2; k++ {
				for l := w - 1; l < w+2; l++ {
					if i == x && j == y && z == k && w == l {

					} else {
						key := strconv.Itoa(i) + "-" + strconv.Itoa(j) + "-" + strconv.Itoa(k) + "-" + strconv.Itoa(l)
						neighbours = append(neighbours, key)
					}
				}
			}
		}

	}
	return
}

func getLowestEncirclingCube(cuby map[string]Cube) (lowX int, highX int, lowY int, highY int, lowZ int, highZ int, lowW int, highW int) {
	lowX = -15
	lowY = -15
	lowZ = -15
	lowW = -15
	highX = 15
	highY = 15
	highZ = 15
	highW = 15
	for _, v := range cuby {
		if v.x < lowX {
			lowX = v.x - 1
		}
		if v.y < lowY {
			lowY = v.y - 1
		}
		if v.z < lowZ {
			lowZ = v.z - 1
		}
		if v.w < lowW {
			lowW = v.w - 1
		}

		if v.x > highX {
			highX = v.x + 1
		}
		if v.y > highY {
			highY = v.y + 1
		}
		if v.z > highZ {
			highZ = v.z + 1
		}
		if v.w > highW {
			highW = v.w + 1
		}
	}
	return
}
