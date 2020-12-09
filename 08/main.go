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

func detectLoop(source Source) (accumulator int) {
	vistitedLine := make(map[int]bool)
	var instructionPointer int
	for {
		if _, notVisited := vistitedLine[instructionPointer]; notVisited {
			fmt.Printf("Visited: %+v \n", vistitedLine)
			return accumulator
		}
		vistitedLine[instructionPointer] = true

		instruction := source.operations[instructionPointer]

		if instruction.operation == "nop" {
			instructionPointer++
			continue
		}
		if instruction.operation == "acc" {
			instructionPointer++
			accumulator += instruction.argument
		}
		if instruction.operation == "jmp" {
			instructionPointer += instruction.argument
		}
	}
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
	accumulator := detectLoop(source)
	fmt.Printf("The accumulator was: %+v", accumulator)
}
