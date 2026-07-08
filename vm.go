package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
)

var addresses []int64
const (
	cmpr = iota + 1
	outputr
	insr
	insp
	addr1
	addr2
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
	operand int64
	ra,rb int64
}
var instructionstack [64*1000]instruction

func bint64(o bool) int64 {
	if o {
		return 1
	}
	return 0
}
func ld(register,number int64) {
	if register == outputr {
		fmt.Println(number)
	}
	addresses[register] = number
}
func mov(registera,registerb int64) {
	if registera == outputr {
		fmt.Println(addresses[registerb])
	}
	addresses[registera] = addresses[registerb]
}
func add(registera,registerb int64) {
	addresses[registera] += addresses[registerb]
}
func mul(registera,registerb int64) {
	addresses[registera] *= addresses[registerb]
}
func div(registera,registerb int64) {
	addresses[registera] /= addresses[registerb]
}
func sub(registera,registerb int64) {
	addresses[registera] -= addresses[registerb]
}

func xor(registera,registerb int64) {
	addresses[registera] ^= addresses[registerb]
}
func and(registera,registerb int64) {
	addresses[registera] &= addresses[registerb]
}
func or(registera,registerb int64) {
	addresses[registera] |= addresses[registerb]
}
func not(registera,registerb int64) {
	addresses[registera] = ^addresses[registera]
}
func cmp(registera,registerb int64) {
	addresses[cmpr] = bint64((addresses[registera] == addresses[registerb])) | (bint64((addresses[registera] < addresses[registerb])) << 1) | (bint64((addresses[registera] > addresses[registerb])) << 2)
}
func movne(registera,registerb int64) {
	if addresses[cmpr]&0b1 != 1 {
		mov(registera,registerb)
	}
}
func move(registera,registerb int64) {
	fmt.Println(addresses[registerb])
	if addresses[cmpr]&0b1 == 1 {
		mov(registera,registerb)
	}
}
func movl(registera,registerb int64) {
	if addresses[cmpr]&0b01 == 1 {
		mov(registera,registerb)
	}
}
func movg(registera,registerb int64) {
	if addresses[cmpr]&0b001 == 1 {
		mov(registera,registerb)
	}
}
func main() {
	addresses = make([]int64, r10)
	var instructions = []func(int64,int64){
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
		movne,
		move,
		movl,
		movg,
	}
	var current instruction
	var humaninstruction string
	var err error
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(addresses[insp], ":")
 		humaninstruction,err = stdin.ReadString('\n')
		humaninstruction = humaninstruction[:len(humaninstruction)-1]
		if len(humaninstruction) < 3 || len(humaninstruction) > 5 {
			fmt.Println("instruction length wrong",humaninstruction)
			continue
		}
		current.operand ,err = strconv.ParseInt(string(humaninstruction[0]),16,32);
		if err != nil {
			fmt.Println(err);
			continue
		}
		current.ra ,err = strconv.ParseInt(string(humaninstruction[1:3]),16,32);
		if err != nil {
			fmt.Println(err);
			continue
		}
		current.rb ,err = strconv.ParseInt(string(humaninstruction[3:5]),16,32);
		if err != nil {
			fmt.Println(err);
			continue
		}
		for i := addresses[insp] ; instructionstack[i].operand != 0; i++ {
			addresses[insp] = i;
		}
		instructionstack[addresses[insp]] = current
		for instructionstack[addresses[insp]].operand != 0 {
			addresses[insr] = instructionstack[addresses[insp]].operand
			addresses[addr1] = instructionstack[addresses[insp]].ra
			addresses[addr2] = instructionstack[addresses[insp]].rb

			instructions[addresses[insr]-1](addresses[addr1],addresses[addr2])
			addresses[insp]++
		}
	}

}
