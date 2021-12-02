package main

import (
	"strings"
	"testing"
)

const testdata = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

func Test_main(t *testing.T) {
	var result1, result2 int

	for _, line := range strings.Split(testdata, "\n") {
		fields := strings.Split(line, " ")
		password := parseFields(fields)
		result1 = result1 + checkPassword(password)
		result2 = result2 + checkPassword2(password)
	}
	if result1 != 2 {
		t.Errorf("wrong result = %v, want %v", result1, 2)
	}
	if result2 != 1 {
		t.Errorf("wrong result = %v, want %v", result2, 1)
	}
}
