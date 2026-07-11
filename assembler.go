package main

import (
	"slices"
	//"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		"exit": 0,
		"ld":    1,
		"mov":   2,
		"add":   3,
		"mul":   4,
		"div":   5,
		"sub":   6,
		"not":   7,
		"and":   8,
		"or":    9,
		"xor":   10,
		"cmp":   11,
		"movne": 12,
		"move":  13,
		"movl":  14,
		"movg":  15,
		"dump":  16,
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
		fmt.Fprintln(os.Stderr, "cant open file", err)
	}
	stringedinput := string(input)
	stringedinput = strings.ReplaceAll(stringedinput, "\t", "")
	lines := strings.Split(stringedinput, "\n")
	lines = slices.DeleteFunc(lines, func(s string) bool { return s == "" })
	var output string
	var ra, rb int
	var ok bool
	for i := range lines {
		splitstring := strings.Split(lines[i], " ")
		if len(splitstring) != 2 {
			fmt.Fprintln(os.Stderr, "no operand on line", i, "(\"", lines[i], "\")")
			continue
		}
		operand := instructions[splitstring[0]]
		registersstring := strings.Split(splitstring[1], ",")
		ra, err = strconv.Atoi(registersstring[0])
		if err != nil {
			ra, ok = registers[registersstring[0]]
			if !ok {
				fmt.Fprintln(os.Stderr, "error ", registersstring[0], "not a known register nor is it a number", "(", err, ")")
			}

		}
		rb, err = strconv.Atoi(registersstring[1])
		if err != nil {
			rb, ok = registers[registersstring[1]]
			if !ok {
				fmt.Fprintln(os.Stderr, "error ", registersstring[1], "not a known register nor is it a number", "(", err, ")")
			}
		}
		output += paddparseh(int64(operand), 2) + paddparseh(int64(ra), 2) + paddparseh(int64(rb), 2) + "\n"

	}

	fmt.Print(output)
}
