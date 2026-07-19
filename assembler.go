package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"slices"
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
func getvalue(looked string, labels, registers map[string]int) (int, error) {
	var returned int
	var err error
	var ok bool
	returned, err = strconv.Atoi(looked)
	if err != nil {
		returned, ok = registers[looked]
		if !ok {
			returned, ok = labels[looked]
			if !ok {
				return 0, errors.New(fmt.Sprintf("error \"%s\" not a known register nor is it a number (%v) nor is it a label",looked,err))
			}
		}

	}
	return returned, nil
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: %file(assembly)")
		os.Exit(0)
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
		"null":0,
		"cmpr":1,
		"outputr":2,
		"insr":3,
		"insp":4,
		"addr1":5,
		"addr2":6,
		"intr":7,
		"r1":8,
		"r2":9,
		"r3":10,
		"r4":11,
		"r5":12,
		"r6":13,
		"r7":14,
		"r8":15,
		"r9":16,
		"r10":17,

	}
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "cant open file", err)
		os.Exit(-1)
	}
	var collective error
	label := regexp.MustCompile(`.*\:`)
	var labels map[string]int = make(map[string]int)
	stringedinput := string(input)
	stringedinput = strings.ReplaceAll(stringedinput, "\t", "")
	lines := strings.Split(stringedinput, "\n")
	lines = slices.DeleteFunc(lines, func(s string) bool { return s == "" || regexp.MustCompile(`#.*`).Match([]byte(s)) })
	for i := range lines {
		if label.Match([]byte(lines[i])) {
			_, ok := registers[lines[i]]
			if ok {
				collective = errors.Join(collective, errors.New("cannot make label named "+lines[i]+" becuase it is a register"))
			} else {
				labels[lines[i][:len(lines[i])-1]] = i
			}
		}
	}
	lines = slices.DeleteFunc(lines, func(s string) bool { return label.Match([]byte(s)) })
	var output strings.Builder
	var ra, rb int
	for i := range lines {
		splitstring := strings.Split(lines[i], " ")
		if len(splitstring) != 2 {
			errstring := fmt.Sprintf("no operand on line %d (\"%s\")",i,lines[i])
			collective = errors.Join(collective, errors.New(errstring))
			continue
		}
		operand := instructions[splitstring[0]]
		registersstring := strings.Split(splitstring[1], ",")
		if len(registersstring) != 2 {
			collective = errors.Join(collective, errors.New("not enough args to instruction"))
			continue
		}
		ra, err = getvalue(registersstring[0], labels, registers)
		collective = errors.Join(collective, err)
		rb, err = getvalue(registersstring[1], labels, registers)
		collective = errors.Join(collective, err)
		output.WriteString(paddparseh(int64(operand), 2) + paddparseh(int64(ra), 2) + paddparseh(int64(rb), 2) + "\n")

	}
	if collective == nil {
		fmt.Print(output.String())
	} else {
		fmt.Fprintln(os.Stderr, collective)
	}
}
