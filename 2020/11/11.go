package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Seat struct {
	x       int
	y       int
	val     string
	adjList []string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ylocal := 0
	seats := make(map[string]Seat)
	xlocal := 0
	for scanner.Scan() {

		line := scanner.Text()
		xlocal = 0
		for _, a := range line {

			key := strconv.Itoa(xlocal) + "-" + strconv.Itoa(ylocal)

			key1 := strconv.Itoa(xlocal-1) + "-" + strconv.Itoa(ylocal-1)
			key2 := strconv.Itoa(xlocal-1) + "-" + strconv.Itoa(ylocal)
			key3 := strconv.Itoa(xlocal-1) + "-" + strconv.Itoa(ylocal+1)
			key4 := strconv.Itoa(xlocal) + "-" + strconv.Itoa(ylocal-1)
			key5 := strconv.Itoa(xlocal) + "-" + strconv.Itoa(ylocal+1)
			key6 := strconv.Itoa(xlocal+1) + "-" + strconv.Itoa(ylocal-1)
			key7 := strconv.Itoa(xlocal+1) + "-" + strconv.Itoa(ylocal)
			key8 := strconv.Itoa(xlocal+1) + "-" + strconv.Itoa(ylocal+1)

			adjacenyList := []string{key1, key2, key3, key4, key5, key6, key7, key8}

			a := Seat{x: xlocal, y: ylocal, val: string(a), adjList: adjacenyList}
			seats[key] = a
			xlocal = xlocal + 1

			//fmt.Println(xlocal, ylocal)
		}
		ylocal = ylocal + 1
	}
	//fmt.Println(seats, xlocal, ylocal)

	chaos := 999

	for chaos != 0 {
		seats, chaos = round(seats, ylocal, xlocal)

	}
	seatsTaken := 0
	for _, a := range seats {
		if a.val == "#" {
			seatsTaken++
		}
	}
	fmt.Println("Seats occupied", seatsTaken)
	printSeats(seats, xlocal, ylocal)

	//	printSeats(seats, xlocal, ylocal)

}

func printSeats(seats map[string]Seat, xMax int, yMax int) {

	fmt.Println(xMax, yMax)
	for i := 0; i < xMax; i++ {
		for j := 0; j < yMax; j++ {
			key := strconv.Itoa(j) + "-" + strconv.Itoa(i)
			fmt.Print(seats[key].val)
		}
		fmt.Println("")
	}
}

func round(seats map[string]Seat, xMax int, yMax int) (result map[string]Seat, changes int) {
	result = make(map[string]Seat)

	fmt.Println(xMax, yMax)
	changes = 0
	for i := 0; i < xMax; i++ {
		for j := 0; j < yMax; j++ {
			key := strconv.Itoa(j) + "-" + strconv.Itoa(i)
			visi := getVisibleSeats(key, seats)
			//	fmt.Println(visi)
			switch seats[key].val {
			case "L":
				canSit := true

				for _, k := range visi {
					if _, ok := seats[k]; ok {
						if seats[k].val == "#" {
							canSit = false
						}
					}
				}

				if canSit {
					changes++
					seatChange := seats[key]
					seatChange.val = "#"
					result[key] = seatChange
				} else {
					result[key] = seats[key]
				}
			case "#":
				seatsBusy := 0

				for _, k := range visi {
					if _, ok := seats[k]; ok {
						if seats[k].val == "#" {
							seatsBusy++
						}
					}
				}

				if 5 <= seatsBusy {
					changes++
					seatChange := seats[key]
					seatChange.val = "L"
					result[key] = seatChange
				} else {
					result[key] = seats[key]
				}

			default:
				result[key] = seats[key]
			}
		}
	}
	fmt.Println("cahnges", changes)
	return
}

func getVisibleSeats(key string, seats map[string]Seat) (lookers []string) {
	distance := 1
	foundKey1 := false
	for !foundKey1 {
		key1 := strconv.Itoa(seats[key].x-distance) + "-" + strconv.Itoa(seats[key].y-distance)
		if _, ok := seats[key1]; ok {
			if seats[key1].val != "." {
				foundKey1 = true
				lookers = append(lookers, key1)
			}
		} else {
			foundKey1 = true
		}
		distance = distance + 1
	}

	distance = 1
	foundKey2 := false
	for !foundKey2 {
		key2 := strconv.Itoa(seats[key].x-distance) + "-" + strconv.Itoa(seats[key].y)
		if _, ok := seats[key2]; ok {
			if seats[key2].val != "." {
				foundKey2 = true
				lookers = append(lookers, key2)
			}
		} else {
			foundKey2 = true
		}
		distance = distance + 1
	}

	distance = 1
	foundKey3 := false
	for !foundKey3 {
		key3 := strconv.Itoa(seats[key].x-distance) + "-" + strconv.Itoa(seats[key].y+distance)
		if _, ok := seats[key3]; ok {
			if seats[key3].val != "." {
				foundKey3 = true
				lookers = append(lookers, key3)
			}
		} else {
			foundKey3 = true
		}
		distance = distance + 1
	}

	distance = 1
	foundKey4 := false
	for !foundKey4 {
		key4 := strconv.Itoa(seats[key].x) + "-" + strconv.Itoa(seats[key].y-distance)
		if _, ok := seats[key4]; ok {
			if seats[key4].val != "." {
				foundKey4 = true
				lookers = append(lookers, key4)
			}
		} else {
			foundKey4 = true
		}
		distance = distance + 1
	}

	distance = 1
	foundKey5 := false
	for !foundKey5 {
		key5 := strconv.Itoa(seats[key].x) + "-" + strconv.Itoa(seats[key].y+distance)
		if _, ok := seats[key5]; ok {
			if seats[key5].val != "." {
				foundKey5 = true
				lookers = append(lookers, key5)
			}
		} else {
			foundKey5 = true
		}
		distance = distance + 1
	}

	distance = 1
	foundKey6 := false
	for !foundKey6 {
		key6 := strconv.Itoa(seats[key].x+distance) + "-" + strconv.Itoa(seats[key].y-distance)
		if _, ok := seats[key6]; ok {
			if seats[key6].val != "." {
				foundKey6 = true
				lookers = append(lookers, key6)
			}
		} else {
			foundKey6 = true
		}
		distance = distance + 1
	}

	distance = 1
	foundKey7 := false
	for !foundKey7 {
		key7 := strconv.Itoa(seats[key].x+distance) + "-" + strconv.Itoa(seats[key].y)
		if _, ok := seats[key7]; ok {
			if seats[key7].val != "." {
				foundKey7 = true
				lookers = append(lookers, key7)
			}
		} else {
			foundKey7 = true
		}
		distance = distance + 1
	}

	distance = 1
	foundKey8 := false
	for !foundKey8 {
		key8 := strconv.Itoa(seats[key].x+distance) + "-" + strconv.Itoa(seats[key].y+distance)
		if _, ok := seats[key8]; ok {
			if seats[key8].val != "." {
				foundKey8 = true
				lookers = append(lookers, key8)
			}
		} else {
			foundKey8 = true
		}
		distance = distance + 1
	}
	return

}
