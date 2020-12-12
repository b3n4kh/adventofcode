package main

import (
	"fmt"
	"strings"
	"testing"
)

const testdata = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func Test_main(t *testing.T) {
	for _, line := range strings.Split(testdata, "\n") {
		fmt.Println(line)
	}

}
