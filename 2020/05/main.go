package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Seat destination
type Seat struct {
	row byte
	col byte
}

func bitStringToBytes(s string) byte {
	var b byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		b |= (c - '0') << uint(7-i&7)
	}
	return b
}

func parseBoardingPass(line string) (seat Seat) {
	row := strings.ReplaceAll(line[0:7], "F", "0")
	row = strings.ReplaceAll(row, "B", "1")
	col := strings.ReplaceAll(line[7:10], "L", "0")
	col = strings.ReplaceAll(col, "R", "1")

	shiftedRow := bitStringToBytes(row) >> 1
	shiftedCol := bitStringToBytes(col) >> 5

	return Seat{shiftedRow, shiftedCol}
}

func (seat Seat) getID() int {
	return int(seat.row)*8 + int(seat.col)
}

func getEmptySeat(seatIDs []int) (emptySeat int) {
	for i, seatID := range seatIDs {
		if seatID+1 != seatIDs[i+1] {
			return seatID + 1
		}
	}
	return 0
}

func main() {
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var line string
	var seats []Seat
	var seatIDs []int
	for scanner.Scan() {
		line = scanner.Text()
		seat := parseBoardingPass(line)
		seats = append(seats, seat)
		seatIDs = append(seatIDs, seat.getID())
	}
	sort.Ints(seatIDs)

	fmt.Printf("Empty Seat ID: %+v\n", getEmptySeat(seatIDs))
	fmt.Printf("\n\nHighest Seat ID: %+v", seatIDs[len(seatIDs)-1:])
}
