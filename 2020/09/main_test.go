package main

import (
	"sort"
	"strconv"
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
	preamble := 5

	var numbers []int

	var currNumber, result int

	for _, line := range strings.Split(testdata, "\n") {
		currNumber, _ = strconv.Atoi(line)

		if len(numbers) <= preamble {
			numbers = append(numbers, currNumber)
			continue
		}

		numbers = leftRotation(numbers)
		numbers = append(numbers, currNumber)

		result = checkPartOne(numbers, currNumber)
		if result != 0 {
			break
		}
	}
	if result != 127 {
		t.Errorf("The result was: %v wanted %v", result, 127)
	}
}

func Test_part2(t *testing.T) {
	var numbers []int
	partOne := 127

	var currNumber, result, sum int
	var s []int

	for _, line := range strings.Split(testdata, "\n") {
		currNumber, _ = strconv.Atoi(line)
		numbers = append(numbers, currNumber)
	}
	index, count := checkPartTwo(numbers, partOne)

	if index != 0 {
		for i := index; i <= index+count; i++ {
			s = append(s, numbers[i])
			sum += numbers[i]
			if sum == partOne {
				break
			}
		}
		sort.Ints(s)
		result = s[0] + s[len(s)-1]
	}

	if result != 62 {
		t.Errorf("The result was: %v wanted %v", result, 62)
	}
}
