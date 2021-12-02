package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Mask repr of binary mask
type Mask [36]int

// ResultMemory holds memory after Mask
type ResultMemory map[int]int

// Mem line from input
type Mem struct {
	addr  int
	value int
}

func parseLine(line string) (key string, value string) {
	split := strings.Split(line, " = ")
	return split[0], split[1]
}

func parseMask(maskvalue string) (mask Mask) {
	for pos, char := range maskvalue {
		if char == 'X' {
			mask[pos] = -1
			continue
		}
		mask[pos], _ = strconv.Atoi(string(char))
	}
	return mask
}

func parseMem(key string, value string) (mem Mem, memBits Mask) {
	address, _ := strconv.Atoi(key[4 : len(key)-1])
	val, _ := strconv.Atoi(value)
	mem = Mem{address, val}
	bitString := fmt.Sprintf("%b", mem.value)

	memBits = stringToBits(bitString)

	return mem, memBits
}

func stringToBits(bitString string) (bits Mask) {
	// fmt.Printf("%v\n", bitString)

	bitPos := 36 - len(bitString)
	for i := 0; i <= (len(bitString) - 1); i++ {
		bits[bitPos], _ = strconv.Atoi(string(bitString[i]))
		bitPos++
	}
	return bits
}

func updateMemory(mask Mask, memory Mask) (result Mask) {
	// fmt.Printf("mask: %+v\nmem:  %+v\n", mask, memToBits(mem))
	for i, bit := range mask {
		if bit == -1 {
			result[i] = memory[i]
			continue
		}
		result[i] = bit
	}
	return result
}

func (result ResultMemory) updateMemoryPart2(mask Mask, addr Mask, value int) {

	var floatingBits []int
	var resultMask Mask

	// fmt.Printf("mask: %+v\nmem:  %+v\n", mask, memToBits(mem))
	for i, bit := range mask {
		switch bit {
		case 0:
			resultMask[i] = addr[i]
		case 1:
			resultMask[i] = 1
		case -1:
			floatingBits = append(floatingBits, i)
		}
	}
	addressRange := int(math.Pow(2, float64(len(floatingBits))))
	fmtString := "%0" + strconv.Itoa(len(floatingBits)) + "b"
	for i := 0; i < addressRange; i++ {
		fBits := fmt.Sprintf(fmtString, i)
		//fmt.Printf("addrRange: %v, fBits: %+v\n", addressRange, fBits)
		for bitpos, maskpos := range floatingBits {
			b, _ := strconv.Atoi(string(fBits[bitpos]))
			resultMask[maskpos] = b
		}
		result[bitToInt(resultMask)] = value
	}

}

func bitToInt(resultBits Mask) (result int) {
	bits := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(resultBits)), ""), "[]")
	parsed, _ := strconv.ParseInt(bits, 2, 64)
	result = int(parsed)

	return result
}

func main() {
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var mask, res, memBits Mask
	var mem Mem
	resultMemory := make(ResultMemory)
	result2Memory := make(ResultMemory)

	var line string

	for scanner.Scan() {
		line = scanner.Text()
		key, value := parseLine(line)
		if key == "mask" {
			mask = parseMask(value)
			continue
		}
		mem, memBits = parseMem(key, value)
		// fmt.Printf("mem: %+v\n", memBits)

		res = updateMemory(mask, memBits)

		addrBits := stringToBits(fmt.Sprintf("%b", mem.addr))
		result2Memory.updateMemoryPart2(mask, addrBits, mem.value)

		// result2Memory[mem.addr] = parseResult(res2)

		resultMemory[mem.addr] = bitToInt(res)
	}
	// fmt.Printf("resultMemory: %+v\n", resultMemory)
	result := 0
	for _, val := range resultMemory {
		result += val
	}
	fmt.Printf("result: %+v\n", result)

	result = 0
	for _, val := range result2Memory {
		result += val
	}
	fmt.Printf("result2: %+v\n", result)

}
