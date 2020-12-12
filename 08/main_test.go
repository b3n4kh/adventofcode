package main

import (
	"strings"
	"testing"
)

const testdata = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

func Test_main(t *testing.T) {
	var source Source
	var lineCount int
	for _, line := range strings.Split(testdata, "\n") {
		instruction := getInstruction(lineCount, line)
		source.operations = append(source.operations, instruction)
		lineCount++
	}
	accumulator, loop := detectLoop(source)
	if accumulator != 5 && loop {
		t.Errorf("The value in the accumulator was: %v wanted %v", accumulator, 5)
	}
}

func Test_repair(t *testing.T) {
	var source Source
	var lineCount int
	for _, line := range strings.Split(testdata, "\n") {
		instruction := getInstruction(lineCount, line)
		source.operations = append(source.operations, instruction)
		lineCount++
	}
	accumulator := repairCode(source)
	if accumulator != 8 {
		t.Errorf("The value in the accumulator was: %v wanted %v", accumulator, 8)
	}
}
