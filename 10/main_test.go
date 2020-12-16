package main

import (
	"strconv"
	"strings"
	"testing"
)

const testdata = `16
10
15
5
1
11
7
19
6
12
4`

const testdata2 = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func Test_main(t *testing.T) {
	var currNumber int
	var numbers []int

	for _, line := range strings.Split(testdata, "\n") {
		currNumber, _ = strconv.Atoi(line)
		numbers = append(numbers, currNumber)
	}
	oneJolt, threeJolt := getdifferences(numbers)
	result := oneJolt * threeJolt
	if oneJolt != 7 || threeJolt != 5 {
		t.Errorf("The result was: %v + %v = %v wanted %v + %v = %v ", oneJolt, threeJolt, result, 7, 5, 35)
	}
}

func Test_two(t *testing.T) {
	var currNumber int
	var numbers []int

	for _, line := range strings.Split(testdata2, "\n") {
		currNumber, _ = strconv.Atoi(line)
		numbers = append(numbers, currNumber)
	}
	oneJolt, threeJolt := getdifferences(numbers)
	result := oneJolt * threeJolt
	if oneJolt != 22 || threeJolt != 10 {
		t.Errorf("The result was: %v + %v = %v wanted %v + %v = %v ", oneJolt, threeJolt, result, 22, 10, 220)
	}
}
