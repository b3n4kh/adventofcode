package main

import (
	"fmt"
	"strings"
	"testing"
)

const testdata = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

const testdata2 = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

func Test_main(t *testing.T) {
	var mask, res, memBits Mask
	var mem Mem
	resultMemory := make(ResultMemory)
	for _, line := range strings.Split(testdata, "\n") {
		key, value := parseLine(line)
		if key == "mask" {
			mask = parseMask(value)
			continue
		}
		mem, memBits = parseMem(key, value)
		// fmt.Printf("mem: %+v\n", memBits)

		res = updateMemory(mask, memBits)
		resultMemory[mem.addr] = bitToInt(res)
	}

	result := 0
	for _, val := range resultMemory {
		result += val
	}

	if result != 165 {
		t.Errorf("The result was: %v wanted %v ", result, 165)
	}
}

func Test_main2(t *testing.T) {
	var mask Mask
	var mem Mem
	resultMemory := make(ResultMemory)
	for _, line := range strings.Split(testdata2, "\n") {
		key, value := parseLine(line)
		if key == "mask" {
			mask = parseMask(value)
			continue
		}
		mem, _ = parseMem(key, value)
		addrBits := stringToBits(fmt.Sprintf("%b", mem.addr))

		resultMemory.updateMemoryPart2(mask, addrBits, mem.value)
	}

	result := 0
	for _, val := range resultMemory {
		result += val
	}

	if result != 208 {
		t.Errorf("The result was: %v wanted %v ", result, 208)
	}
}
