package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Instruction struct {
	tot  string
	key  string
	pace int
}

type Ship struct {
	x      int
	y      int
	facing int
}

type Waypoint struct {
	x int
	y int
}

func main() {
	var instructions []Instruction
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		line := scanner.Text()
		k := line[:1]
		p, _ := strconv.Atoi(line[1:])
		instr := Instruction{tot: line, key: k, pace: p}
		instructions = append(instructions, instr)
	}
	fmt.Println(instructions)
	ship := Ship{x: 0, y: 0, facing: 90}
	waypoint := Waypoint{x: 10, y: 1}

	for _, instruct := range instructions {
		switch instruct.key {
		case "N":
			waypoint.y = waypoint.y + instruct.pace
		case "S":
			waypoint.y = waypoint.y - instruct.pace
		case "E":
			waypoint.x = waypoint.x + instruct.pace
		case "W":
			waypoint.x = waypoint.x - instruct.pace
		case "L":
			ship.facing = ((ship.facing - instruct.pace) + 360) % 360
			switch instruct.pace {
			case 180:
				waypoint.x = waypoint.x * -1
				waypoint.y = waypoint.y * -1

			case 90:
				temp := waypoint.x
				waypoint.x = waypoint.y * -1
				waypoint.y = temp
			case 270:
				//ship.y = ship.y + instruct.pace
				temp := waypoint.x
				waypoint.x = waypoint.y
				waypoint.y = temp * -1
			}
		case "R":
			ship.facing = (ship.facing + instruct.pace) % 360
			switch instruct.pace {
			case 180:
				waypoint.x = waypoint.x * -1
				waypoint.y = waypoint.y * -1

			case 270:
				temp := waypoint.x
				waypoint.x = waypoint.y * -1
				waypoint.y = temp
			case 90:
				//ship.y = ship.y + instruct.pace
				temp := waypoint.x
				waypoint.x = waypoint.y
				waypoint.y = temp * -1
			}
		case "F":
			ship.x = ship.x + (instruct.pace * waypoint.x)
			ship.y = ship.y + (instruct.pace * waypoint.y)

		default:
		}
		fmt.Println(ship, waypoint)
		if waypoint.x > 0 {
			if waypoint.y > 0 {
				fmt.Println("East", waypoint.x, "north", waypoint.y)
			} else {
				fmt.Println("East", waypoint.x, "south", waypoint.y)
			}
		} else {
			if waypoint.y > 0 {
				fmt.Println("West", waypoint.x, "north", waypoint.y)
			} else {
				fmt.Println("West", waypoint.x, "south", waypoint.y)
			}
		}
	}
	fmt.Println(ship)
}
