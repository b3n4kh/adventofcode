package main

import (
	"bufio"
	"os"
	"strings"
)

// Seat destination
type Seat struct {
	row byte
	col byte
	id  int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseBoardingPass(line string) (seat Seat) {
	row := strings.ReplaceAll(line[0:7], "F", "0")
	row = strings.ReplaceAll(row, "B", "1")
	col := strings.ReplaceAll(line[7:10], "L", "0")
	col = strings.ReplaceAll(col, "R", "1")

	return Seat{row, col}
}

func main() {
	dat, err := os.Open("input.txt")
	check(err)
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var line string
	for scanner.Scan() {
		line = scanner.Text()

	}
}
