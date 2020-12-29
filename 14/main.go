package main

import (
	"bufio"
	"fmt"
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
	memBits = memToBits(mem)

	return mem, memBits
}

func memToBits(mem Mem) (bits Mask) {
	bitString := fmt.Sprintf("%b", mem.value)
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

func parseResult(resultBits Mask) (result int) {
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
		// fmt.Printf("result: %+v\n", res)
		resultMemory[mem.addr] = parseResult(res)
	}
	// fmt.Printf("resultMemory: %+v\n", resultMemory)
	var result int
	for _, val := range resultMemory {
		result += val
	}
	fmt.Printf("result: %+v\n", result)

}
