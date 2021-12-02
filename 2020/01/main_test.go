package main

import (
	"strconv"
	"strings"
	"testing"
)

const testdata = `1721
979
366
299
675
1456`

func Test_main(t *testing.T) {
	var result []int

	for _, line := range strings.Split(testdata, "\n") {
		x, err := strconv.Atoi(line)
		check(err)
		result = append(result, x)
	}
	two, three := find2020(result)

	if two != 514579 {
		t.Errorf("wrong result = %v, want %v", two, 514579)
	}
	if three != 241861950 {
		t.Errorf("wrong result = %v, want %v", three, 241861950)
	}
}
