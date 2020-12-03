package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type flightState struct {
	position int
	trees    int
}

type slope struct {
	right int
	down  int
}

var slopes = []slope{
	{1, 1},
	{3, 1},
	{5, 1},
	{7, 1},
	{1, 2},
}

func isTree(character byte) bool {
	return character == '#'
}

func getLine(line string, position int, stepright int) string {
	if (position + stepright) > len(line) {
		return getLine((line + line), position, stepright)
	}
	return line
}

func flightControl(scanner *bufio.Scanner, slope slope, position int, trees int) int {
	line := getLine(scanner.Text(), position, slope.right)

	if isTree(line[position]) {
		trees++
	}
	for i := 0; i < slope.down; i++ {
		if !scanner.Scan() {
			return trees
		}
	}
	return flightControl(scanner, slope, position+slope.right, trees)
}

func main() {
	result := 1
	for _, slope := range slopes {
		dat, err := ioutil.ReadFile("input.txt")
		check(err)
		scanner := bufio.NewScanner(bytes.NewReader(dat))
		scanner.Scan()
		trees := flightControl(scanner, slope, 0, 0)
		result = result * trees
		fmt.Printf("Slope %v : #Trees: %v\n", slope, trees)
	}
	fmt.Printf("Number Trees overall: %d", result)
}
