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

// Mem line from input
type Mem struct {
	addr  int
	value int
}

type ResultMem sturct {
	value int
}

func parseLine(line string) (key string, value string) {
	split := strings.Split(line, " = ")
	return split[0], split[1]
}

func parseMask(maskvalue string) (mask Mask) {
	for pos, char := range maskvalue {
		if char == 'X' {
			continue
		}
		mask[pos], _ = strconv.Atoi(string(char))
	}
	return mask
}

func parseMem(key string, value string) (mem Mem) {
	address, _ := strconv.Atoi(key[4 : len(key)-1])
	val, _ := strconv.Atoi(value)
	mem = Mem{address, val}
	return mem
}

func updateMemory(mask Mask, mem Mem) {
	//strconv.Atoi(line)
}

func main() {
	dat, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var mask Mask
	var mem Mem
	resultMemory = []Mem

	var line string

	for scanner.Scan() {
		line = scanner.Text()
		key, value := parseLine(line)
		if key == "mask" {
			mask = parseMask(value)
			continue
		}
		mem = parseMem(key, value)
		resultMemory[mem.addr] = updateMemory(mask, mem)
	}
	fmt.Printf("Result1: %v", result)
}
