package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Source struct {
	operations []Instruction
}

type Instruction struct {
	line      int
	operation string
	argument  int
}

func getInstruction(lineCount int, line string) (instruction Instruction) {
	data := strings.Split(line, " ")
	operation := data[0]
	if operation != "nop" && operation != "acc" && operation != "jmp" {
		panic("Could not parse line")
	}
	argument, err := strconv.Atoi(data[1])
	if err != nil {
		panic("Could not parse line")
	}

	return Instruction{lineCount, operation, argument}
}

func detectLoop(source Source) (accumulator int, loop bool) {
	vistitedLine := make(map[int]bool)
	var instructionPointer int
	for {

		if len(source.operations)-1 < instructionPointer {
			return accumulator, false
		}
		if _, notVisited := vistitedLine[instructionPointer]; notVisited {
			return accumulator, true
		}

		vistitedLine[instructionPointer] = true

		instruction := source.operations[instructionPointer]

		switch instruction.operation {
		case "nop":
			instructionPointer++
		case "acc":
			instructionPointer++
			accumulator += instruction.argument
		case "jmp":
			instructionPointer += instruction.argument
		}
	}
}

func patchSource(source Source, instructionPointer int) (patchedSource Source) {
	if len(source.operations) <= instructionPointer {
		return source
	}
	patchedSource = source
	patchedSource.operations = make([]Instruction, len(source.operations))
	copy(patchedSource.operations, source.operations)
	instruction := patchedSource.operations[instructionPointer]
	switch instruction.operation {
	case "nop":
		patchedSource.operations[instructionPointer].operation = "jmp"
		return patchedSource
	case "jmp":
		patchedSource.operations[instructionPointer].operation = "nop"
		return patchedSource
	}
	instructionPointer++
	return patchSource(patchedSource, instructionPointer)
}

func repairCode(source Source) (accumulator int) {
	var loop bool
	for currentInstruction := 0; currentInstruction <= len(source.operations); currentInstruction++ {
		patchedSource := patchSource(source, currentInstruction)

		accumulator, loop = detectLoop(patchedSource)
		//fmt.Printf("Accumulator: %+v loop: %+v\n\n", accumulator, loop)

		if !loop {
			return accumulator
		}
	}
	return 0
}

func main() {
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var line string
	var source Source
	var lineCount int
	for scanner.Scan() {
		line = scanner.Text()
		instruction := getInstruction(lineCount, line)
		source.operations = append(source.operations, instruction)
		lineCount++
	}
	accumulator, loop := detectLoop(source)
	if loop {
		fmt.Printf("The accumulator was: %+v\n", accumulator)
	}

	accumulator = repairCode(source)
	fmt.Printf("The repaired Code accumulator was: %+v\n", accumulator)

}
