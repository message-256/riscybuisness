package main

import (
	"fmt"
)

var addresses []int

const (
	cmpr = iota + 1
	outputr
	constr
	insr
	insp
	r1
	r2
	r3
	r4
	r5
	r6
	r7
	r8
	r9
	r10
)

type instruction struct {
	operand int
	number  []int
}

func bint(o bool) int {
	if o {
		return 1
	}
	return 0
}
func ld(number []int) {
	addresses[constr] = int(number[0])
}
func mov(number []int) {
	if number[0] == outputr {
		fmt.Println(addresses[number[1]])
	}
	addresses[number[0]] = addresses[number[1]]
}
func add(number []int) {
	addresses[number[0]] += addresses[number[1]]
}
func mul(number []int) {
	addresses[number[0]] *= addresses[number[1]]
}
func div(number []int) {
	addresses[number[0]] /= addresses[number[1]]
}
func sub(number []int) {
	addresses[number[0]] -= addresses[number[1]]
}

func xor(number []int) {
	addresses[number[0]] ^= addresses[number[1]]
}
func and(number []int) {
	addresses[number[0]] &= addresses[number[1]]
}
func or(number []int) {
	addresses[number[0]] |= addresses[number[1]]
}
func not(number []int) {
	addresses[number[0]] = ^addresses[number[0]]
}
func cmp(number []int) {
	addresses[cmpr] = bint((addresses[number[0]] == addresses[number[1]])) | (bint((addresses[number[0]] < addresses[number[1]])) << 1) | (bint((addresses[number[0]] > addresses[number[1]])) << 2)
}
func jmp(number []int) {
	addresses[insp] = addresses[number[0]]
}
func jne(number []int) {
	if addresses[cmpr]&0b1 != 1 {
		addresses[insp] = addresses[number[0]]
	}
}
func je(number []int) {
	fmt.Println(addresses[number[0]])
	if addresses[cmpr]&0b1 == 1 {
		addresses[insp] = addresses[number[0]]
	}
}
func jl(number []int) {
	if addresses[cmpr]&0b01 == 1 {
		addresses[insp] = addresses[number[0]]
	}
}
func jg(number []int) {
	if addresses[cmpr]&0b001 == 1 {
		addresses[insp] = addresses[number[0]]
	}
}
func main() {
	addresses = make([]int, r10)
	var instructions = []func([]int){
		ld,
		mov,
		add,
		mul,
		div,
		sub,
		not,
		and,
		or,
		xor,
		cmp,

		jmp,
		je,
		jne,
		jl,
		jg,
	}
	var program []instruction
	var current instruction
	var humaninstruction string
	var hexstuff = map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'a': 10,
		'b': 11,
		'c': 12,
		'd': 13,
		'e': 14,
		'f': 15,
	}
	for {
		fmt.Print(addresses[insp], ":")
		fmt.Scanf("%s", &humaninstruction)
		if len(humaninstruction) < 3 || len(humaninstruction) > 5 {
			fmt.Println("instruction length wrong")
			continue
		}
		current.operand = hexstuff[humaninstruction[0]]
		current.number = make([]int, 1)
		current.number[0] = (hexstuff[humaninstruction[1]] << 4) | hexstuff[humaninstruction[2]]
		if len(humaninstruction) == 5 {
			current.number = append(current.number, (hexstuff[humaninstruction[3]]<<4)|hexstuff[humaninstruction[4]])
		}
		fmt.Println(current)
		program = append(program, current)
		addresses[insp] = len(program) - 1
		for addresses[insp] < len(program) {
			fmt.Println("infinite loop", addresses[insp])
			addresses[insr] = program[addresses[insp]].operand
			instructions[program[addresses[insp]].operand](program[addresses[insp]].number)
			addresses[insp]++
		}
	}

}
