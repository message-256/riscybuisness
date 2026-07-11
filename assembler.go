package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"errors"
)

func paddparseh(in int64, bit int) string {
	parsed := strconv.FormatInt(in, 16)
	output := parsed
	for i := 0; i < bit-len(parsed); i++ {
		output = "0" + output
	}
	return output
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: %file(assembly)")
	}
	var instructions = map[string]int{
		"exit":  0,
		"ld":    1,
		"mov":   2,
		"add":   3,
		"mul":   4,
		"div":   5,
		"sub":   6,
		"not":   7,
		"shl":   8,
		"shr":   9,
		"and":   10,
		"or":    11,
		"xor":   12,
		"cmp":   13,
		"movne": 14,
		"move":  15,
		"movl":  16,
		"movg":  17,
	}
	var registers = map[string]int{
		"cmpr":    0,
		"outputr": 1,
		"insr":    2,
		"insp":    3,
		"addr1":   4,
		"addr2":   5,
		"intr":    6,
		"r1":      7,
		"r2":      8,
		"r3":      9,
		"r4":      10,
		"r5":      11,
		"r6":      12,
		"r7":      13,
		"r8":      14,
		"r9":      15,
		"r10":     16,
	}
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "cant open file", err)
		os.Exit(-1)
	}
	stringedinput := string(input)
	stringedinput = strings.ReplaceAll(stringedinput, "\t", "")
	lines := strings.Split(stringedinput, "\n")
	lines = slices.DeleteFunc(lines, func(s string) bool { return s == "" })
	var output string
	var ra, rb int
	var ok bool
	var collective error
	for i := range lines {
		splitstring := strings.Split(lines[i], " ")
		if len(splitstring) != 2 {
			errstring := fmt.Sprintln( "no operand on line", i, "(\"", lines[i], "\")")
			collective = errors.Join(collective,errors.New(errstring))
			continue
		}
		operand := instructions[splitstring[0]]
		registersstring := strings.Split(splitstring[1], ",")
		ra, err = strconv.Atoi(registersstring[0])
		if err != nil {
			ra, ok = registers[registersstring[0]]
			if !ok {
				errstring := fmt.Sprintln( "error ", registersstring[0], "not a known register nor is it a number", "(", err, ")")
				collective = errors.Join(collective,errors.New(errstring))
			}

		}
		rb, err = strconv.Atoi(registersstring[1])
		if err != nil {
			rb, ok = registers[registersstring[1]]
			if !ok {
				errstring := fmt.Sprintln( "error ", registersstring[1], "not a known register nor is it a number", "(", err, ")")
				collective = errors.Join(collective,errors.New(errstring))
			}
		}
		output += paddparseh(int64(operand), 2) + paddparseh(int64(ra), 2) + paddparseh(int64(rb), 2) + "\n"

	}
	if collective == nil{
		fmt.Print(output)
	} else {
		fmt.Fprintln(os.Stderr,collective)
	}
}
