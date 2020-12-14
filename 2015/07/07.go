package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Signal struct {
	input  string
	value  uint16
	line   string
	output string
	ok     bool
}

type OrGate struct {
	inputA string
	inputB string
	output string
}

type AndGate struct {
	inputA string
	inputB string
	output string
}

type LShiftGate struct {
	inputA string
	val    uint16
	output string
}
type RShiftGate struct {
	inputA string
	val    uint16
	output string
}

type NotGate struct {
	inputA string
	output string
}

func main() {
	signalss := make(map[string]Signal)
	var signals []Signal
	var orGates []OrGate
	var andGates []AndGate
	var lShifts []LShiftGate
	var rShifts []RShiftGate
	var notGates []NotGate

	scanner := bufio.NewScanner(os.Stdin)
	var instructions []string
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, line)
		args := strings.Split(line, " ")

		if len(args) == 3 {
			done := true

			val, err := strconv.ParseUint(args[0], 10, 16)
			if err != nil {
				done = false
			}
			signal := Signal{input: args[0], value: uint16(val), line: line, ok: done, output: args[2]}
			signals = append(signals, signal)
		}

		if len(args) == 4 {
			g := NotGate{inputA: args[1], output: args[3]}
			notGates = append(notGates, g)
		}

		if len(args) == 5 {

			if args[1] == "AND" {
				g := AndGate{inputA: args[0], inputB: args[2], output: args[4]}
				andGates = append(andGates, g)

			}
			if args[1] == "OR" {
				g := OrGate{inputA: args[0], inputB: args[2], output: args[4]}
				orGates = append(orGates, g)

			}
			if args[1] == "LSHIFT" {
				shiftStep, _ := strconv.ParseUint(args[2], 10, 16)

				g := LShiftGate{inputA: args[0], val: uint16(shiftStep), output: args[4]}
				lShifts = append(lShifts, g)

			}
			if args[1] == "RSHIFT" {
				shiftStep, _ := strconv.ParseUint(args[2], 10, 16)

				g := RShiftGate{inputA: args[0], val: uint16(shiftStep), output: args[4]}
				rShifts = append(rShifts, g)
			}
		}
	}
	fmt.Println(signals)

	i := 0

	for {
		var doneWithSignals []Signal
		var doneWithAnds []AndGate
		var doneWithOrs []OrGate
		var doneWithRShifts []RShiftGate
		var doneWithLShifts []LShiftGate
		var doneWithNots []NotGate

		for _, sig := range signals {

			if sig.ok {
				signalss[sig.output] = sig
			} else {
				if _, ok := signalss[sig.input]; ok {
					s := Signal{value: signalss[sig.input].value, ok: true, input: sig.input}
					signalss[sig.output] = s
				} else {
					doneWithSignals = append(doneWithSignals, sig)

				}
			}
		}
		for _, g := range andGates {
			var ll uint16

			added := false
			val, err := strconv.ParseUint(g.inputA, 10, 16)
			if err == nil {
				ll = uint16(val)
				if b, ok := signalss[g.inputB]; ok {

					output := Signal{value: ll & b.value, output: g.output}
					signalss[g.output] = output
					added = true
				}
			} else {
				if a, ok := signalss[g.inputA]; ok {
					if b, ok := signalss[g.inputB]; ok {

						output := Signal{value: a.value & b.value, output: g.output}
						signalss[g.output] = output
						added = true
					}
				}
			}

			if !added {
				doneWithAnds = append(doneWithAnds, g)
			}
		}

		for _, g := range orGates {
			added := false

			if a, ok := signalss[g.inputA]; ok {
				if b, ok := signalss[g.inputB]; ok {
					output := Signal{value: a.value | b.value, output: g.output}
					signalss[g.output] = output
					added = true
				}
			}
			if !added {
				doneWithOrs = append(doneWithOrs, g)
			}
		}

		for _, g := range lShifts {
			added := false

			if a, ok := signalss[g.inputA]; ok {
				output := Signal{value: a.value << g.val, output: g.output}
				signalss[g.output] = output
				added = true
			}
			if !added {
				doneWithLShifts = append(doneWithLShifts, g)
			}
		}

		for _, g := range rShifts {
			added := false
			if a, ok := signalss[g.inputA]; ok {
				output := Signal{value: a.value >> g.val, output: g.output}
				signalss[g.output] = output
				added = true
			}
			if !added {
				doneWithRShifts = append(doneWithRShifts, g)
			}
		}
		for _, g := range notGates {
			added := false

			if a, ok := signalss[g.inputA]; ok {
				output := Signal{value: ^a.value, output: g.output}
				signalss[g.output] = output
				added = true
			}
			if !added {
				doneWithNots = append(doneWithNots, g)
			}
		}
		if i > 200 {
			break
		}
		i++
		signals = doneWithSignals
		andGates = doneWithAnds
		orGates = doneWithOrs
		lShifts = doneWithLShifts
		rShifts = doneWithRShifts
		notGates = doneWithNots
		fmt.Println(len(signals), len(andGates),
			len(orGates), len(lShifts),
			len(rShifts), len(notGates))

	}

	for k, v := range signalss {
		if v.value > 0 {
			fmt.Println(k, ":  ", v.value)
		}
	}

}
