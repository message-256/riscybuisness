package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"errors"
	"strings"
	"slices"
)

var addresses []int64 = make([]int64, 0xfffF)

const (
	cmpr = iota
	outputr
	insr
	insp
	addr1
	addr2
	intr
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
	ra, rb  int64
}

var instructionstack [64 * 1000]instruction
var instructions = []func(int64, int64){
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
		dump,
}
var registers = []string {
		"cmpr",
		"outputr",
		"insr",
		"insp",
		"addr1",
		"addr2",
		"intr",
		"r1",
		"r2",
		"r3",
		"r4",
		"r5",
		"r6",
		"r7",
		"r8",
		"r9",
		"r10",


	}
func bint64(o bool) int64 {
	if o {
		return 1
	}
	return 0
}
func dump(r1,r2 int64){
	
	for i := cmpr; i< r10; i++ {
		if i == addr1 || i == addr2 {
			fmt.Println(registers[i],"=",registers[addresses[i]])

		} else {
			fmt.Println(registers[i],"=",addresses[i])
		}
	}
	fmt.Println("done")
}
func ld(register, number int64) {
	if register == outputr {
		fmt.Println(number)
	}
	addresses[register] = number
}
func mov(registera, registerb int64) {
	ld(registera,addresses[registerb])
}
func add(registera, registerb int64) {
	addresses[registera] += addresses[registerb]
}
func mul(registera, registerb int64) {
	addresses[registera] *= addresses[registerb]
}
func div(registera, registerb int64) {
	addresses[registera] /= addresses[registerb]
}
func sub(registera, registerb int64) {
	addresses[registera] -= addresses[registerb]
}

func xor(registera, registerb int64) {
	addresses[registera] ^= addresses[registerb]
}
func and(registera, registerb int64) {
	addresses[registera] &= addresses[registerb]
}
func or(registera, registerb int64) {
	addresses[registera] |= addresses[registerb]
}
func not(registera, registerb int64) {
	addresses[registera] = ^addresses[registera]
}
func cmp(registera, registerb int64) {
	addresses[cmpr] = bint64((addresses[registera] == addresses[registerb])) | (bint64((addresses[registera] < addresses[registerb])) << 1) | (bint64((addresses[registera] > addresses[registerb])) << 2)
	fmt.Println(addresses[cmpr])
}
func movne(registera, registerb int64) {
	if addresses[cmpr]&0b1 != 1 {
		mov(registera, registerb)
	}
}
func move(registera, registerb int64) {
	fmt.Println(addresses[registerb])
	if addresses[cmpr]&0b1 == 1 {
		mov(registera, registerb)
	}
}
func movl(registera, registerb int64) {
	if addresses[cmpr]&0b10 == 0b10 {
		mov(registera, registerb)
	}
}
func movg(registera, registerb int64) {
	if addresses[cmpr]&0b100 == 0b100 {
		mov(registera, registerb)
	}
}
func parse(line string) (instruction,error) {
	var given error
	var returned instruction
	var err error
	if len(line) != 6 {
		return returned,errors.New("instruction length wrong"+ line+ "(length "+ strconv.Itoa(len(line)) + ") is not 6")
	}
	returned.operand, err = strconv.ParseInt(string(line[0:2]), 16, 32)
	if err != nil {
		given = errors.Join(given,err)
	}
	returned.ra, err = strconv.ParseInt(string(line[2:4]), 16, 32)
	if err != nil {
		given = errors.Join(given,err)
	}
	returned.rb, err = strconv.ParseInt(string(line[4:6]), 16, 32)
	if err != nil {
		given = errors.Join(given,err)
	}
	return returned,given
	
}
func run () {
	for instructionstack[addresses[insp]].operand != 0 && addresses[insp] < int64(len(instructionstack)) {
		addresses[insr] = instructionstack[addresses[insp]].operand
		addresses[addr1] = instructionstack[addresses[insp]].ra
		addresses[addr2] = instructionstack[addresses[insp]].rb
		instructions[addresses[insr]-1](addresses[addr1], addresses[addr2])
		dump(0,0)

		addresses[insp]++
	}

}
func main() {
	if len(os.Args) > 1 {
		bytes ,err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr,err);
			os.Exit(-1);
		}
		lines := strings.Split(string(bytes),"\n")
		if len(lines) > len(instructionstack) {
			fmt.Fprintln(os.Stderr,"file larger than instruction stack")
			os.Exit(-1)
		}
		lines = slices.DeleteFunc(lines,func(s string) bool{return s == ""})

		var newerr error
		for i := range lines {
			instructionstack[i],newerr = parse(lines[i]);
			err = errors.Join(err,newerr)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr,err);
			os.Exit(-1)
		}
		run()
		os.Exit(0)
	}
	
	var current instruction
	var humaninstruction string
	var err error
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(addresses[insp], ":")
		humaninstruction, err = stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err);
			continue
		}
		humaninstruction = strings.ReplaceAll(humaninstruction," ","")
		humaninstruction = humaninstruction[:len(humaninstruction)-1]
		current,err = parse(humaninstruction)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for i := addresses[insp]; instructionstack[i].operand != 0; i++ {
			addresses[insp] = i
		}
		instructionstack[addresses[insp]] = current
		fmt.Println(current)
		run()
	}

}
