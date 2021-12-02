package main

import (
	"bufio"
	"strings"
	"testing"
)

const testdata = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func Test_main(t *testing.T) {
	result := 1
	for _, slope := range slopes {
		scanner := bufio.NewScanner(strings.NewReader(testdata))
		scanner.Scan()
		trees := flightControl(scanner, slope, 0, 0)
		result = result * trees
	}

	if result != 336 {
		t.Errorf("wrong result = %v, want %v", result, 336)
	}
}
