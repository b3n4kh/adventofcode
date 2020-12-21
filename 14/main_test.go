package main

import (
	"strconv"
	"strings"
	"testing"
)

const testdata = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

func Test_main(t *testing.T) {
	var result int
	for _, line := range strings.Split(testdata, "\n") {
		currNumber, _ := strconv.Atoi(line)
		result += currNumber
	}

	if result != 165 {
		t.Errorf("The result was: %v wanted %v ", result, 165)
	}
}
