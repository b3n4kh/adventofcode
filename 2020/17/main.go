package main

import (
	"bufio"
	"fmt"
	"os"
)

type layer [][]int

func cycle() {

}

func main() {
	dat, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var line string
	grid := make(map[int]layer)
	baseLayer := layer{}
	yPos := 0
	for scanner.Scan() {
		line = scanner.Text()
		baseLayer = append(baseLayer, []int{})
		baseLayer[yPos] = make([]int, len(line))
		for i, char := range line {
			if char == '.' {
				baseLayer[yPos][i] = 0
			}
			if char == '#' {
				baseLayer[yPos][i] = 1
			}
		}
		yPos++
	}
	grid[0] = baseLayer
	fmt.Printf("BaseState: %+v\n", grid)

}
