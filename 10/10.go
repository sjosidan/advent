package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type Asteroid struct {
	x int
	y int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ylocal := 0
	asteroids := make(map[string]Asteroid)

	for scanner.Scan() {
		xlocal := 0
		line := scanner.Text()
		for _, a := range line {
			if string([]rune{a}) == "#" {
				key := strconv.Itoa(xlocal) + "-" + strconv.Itoa(ylocal)

				a := Asteroid{x: xlocal, y: ylocal}
				fmt.Println(a)
				asteroids[key] = a
			} else {
			}
			xlocal = xlocal + 1

			//fmt.Println(xlocal, ylocal)
		}
		ylocal = ylocal + 1

	}
	maxAngles := 0
	var champ string
	var champStation map[float64][]Asteroid

	for kz, az := range asteroids {
		//	fmt.Println("Checking ", kz)
		localAngles := make(map[float64][]Asteroid)
		for k, a := range asteroids {
			if kz != k {
				yDelta := float64(az.y - a.y)
				xDelta := float64(az.x - a.x)
				deg := (math.Atan2(-yDelta, -xDelta) * 180 / math.Pi)
				if deg < 0 {
					deg = deg + 360
				}
				deg = deg + 90
				if deg > 360 {
					deg = deg - 360
				}
				//	fmt.Println(angle, xDelta, yDelta)
				//fmt.Println(angle)
				//	deg = deg + 180
				if xDelta == 0 {
					if yDelta > 0 {
						localAngles[0] = append(localAngles[0], a)
					} else {
						localAngles[180] = append(localAngles[180], a)
					}
				} else if yDelta == 0 {
					if xDelta > 0 {
						localAngles[270] = append(localAngles[270], a)
					} else {
						localAngles[90] = append(localAngles[90], a)
					}
				} else {
					if yDelta > 0 {
						localAngles[deg] = append(localAngles[deg], a)
					} else {
						localAngles[deg] = append(localAngles[deg], a)
					}

				}
			} else {
				//	fmt.Println(kz, k)
			}
			//	fmt.Println(xx, yy, len(localAngles))

		}
		if len(localAngles) > maxAngles {
			champ = kz
			maxAngles = len(localAngles)
			champStation = localAngles
		}
	}
	var keys []float64
	for k, a := range champStation {
		keys = append(keys, k)
		fmt.Println(k, a)
	}
	sort.Sort(sort.Float64Slice(keys))
	fmt.Println(champ, keys)
	i := 0
	for _, k := range keys {
		fmt.Println(i, champStation[k])
		i = i + 1
	}

}
