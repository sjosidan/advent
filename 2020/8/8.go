package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Operation struct {
	name  string
	value int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	program := make([]Operation, 0)
	for scanner.Scan() {
		line := scanner.Text()
		n := line[:3]
		v, _ := strconv.Atoi(line[4:])

		program = append(program, Operation{name: n, value: v})
	}
	runningForever := true
	res := 0
	programCopy := make([]Operation, len(program))
	for runningForever {
		for instr, changeMe := range program {
			copy(programCopy, program)
			switch changeMe.name {
			case "jmp":
				programCopy[instr].name = "nop"
			case "nop":
				programCopy[instr].name = "jmp"
			default:
			}
			res, runningForever = runProgram(programCopy)
			if runningForever {
			} else {
				fmt.Println("Finished", res)

				break
			}
		}

	}

}

func runProgram(program []Operation) (result int, ranForver bool) {
	state := 0
	ranInstructions := make(map[int]Operation)
	accumulator := 0

	terminated := false
	for !terminated {
		if state > len(program)-1 {
			result = accumulator
			ranForver = false
			break
		}

		instruction := program[state]

		if _, ok := ranInstructions[state]; ok {
			result = accumulator
			ranForver = true
			terminated = true
		} else {
			switch instruction.name {
			case "acc":
				accumulator = accumulator + instruction.value
				ranInstructions[state] = instruction
				state = state + 1
			case "jmp":
				ranInstructions[state] = instruction
				state = state + instruction.value
			case "nop":
				ranInstructions[state] = instruction
				state = state + 1
			default:
				fmt.Errorf("SOMETHINg WRONG")
			}
		}
	}
	return
}
