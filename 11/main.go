package main

import (
	"bufio"
	"fmt"
	"os"
)

type Seat struct {
	row        int
	col        int
	empty      bool
	neighboars []int
}

type seats map[int]Seat

func isEmpty(char rune) bool {
	if char == 'L' {
		return true
	}
	return false
}

func (allSeats *seats) getSeats(line string, row int) {
	for col, char := range line {
		if string(char) == "." {
			continue
		}
		(*allSeats)[row+1*col+1] = Seat{row, col, isEmpty(char), []int{}}
	}
}

/* func (allSeats *seats) getNeighbors() {
	for index, seat := range allSeats {
		if ((seat.row) == (row-1) && seat.col == col) || ((seat.row) == (row+1) && seat.col == col) {
			neigbors = append(neigbors, seat)
		}
		if ((seat.col) == (col-1) && seat.row == row) || ((seat.col) == (col+1) && seat.row == row) {
			neigbors = append(neigbors, seat)
		}
		if ((seat.row) == (row-1) && (seat.col) == (col-1)) || ((seat.row) == (row+1) && (seat.col) == (col+1)) {
			neigbors = append(neigbors, seat)
		}
		if ((seat.col) == (col-1) && (seat.row) == (row+1)) || ((seat.col) == (col+1) && (seat.row) == (row-1)) {
			neigbors = append(neigbors, seat)
		}
	}
}
*/

func (allSeats seats) printSeats(row int, col int) {
	for i := 1; i < col; i++ {
		for y := 1; y < row; y++ {
			if _, ok := allSeats[i*y]; ok {
				if allSeats[i*y].empty {
					fmt.Printf("L")
				} else {
					fmt.Printf("#")
				}
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func (allSeats seats) getNeighbors(row int, col int) seats {
	size := row * col
	var neighborSeats seats
	neighborSeats = make(seats)
	for i, seat := range allSeats {
		neighborSeats[i] = seat
		currentSeat := neighborSeats[i]

		if i-col >= 0 {
			currentSeat.neighboars = append(currentSeat.neighboars, i-col)
		}
		if i%col != 0 {
			currentSeat.neighboars = append(currentSeat.neighboars, i-1)
		}
		if (i+1)%col != 0 {
			currentSeat.neighboars = append(currentSeat.neighboars, i+1)
		}
		if i+col < size {
			currentSeat.neighboars = append(currentSeat.neighboars, i+col)
		}
		if ((i - col - 1) >= 0) && (i%col != 0) {
			currentSeat.neighboars = append(currentSeat.neighboars, i-row-1)
		}
		if ((i - col + 1) >= 0) && ((i+1)%col != 0) {
			currentSeat.neighboars = append(currentSeat.neighboars, i-col+1)
		}
		if ((i + col - 1) < size) && (i%col != 0) {
			currentSeat.neighboars = append(currentSeat.neighboars, i+col-1)
		}
		if ((i + col + 1) < size) && ((i+1)%col != 0) {
			currentSeat.neighboars = append(currentSeat.neighboars, i+col+1)
		}
		neighborSeats[i] = currentSeat

	}
	//fmt.Printf("%+v\n", neighborSeats)

	return neighborSeats
}

func changeSeats(allSeats seats) (newSeats seats) {
	newSeats = make(seats)
	var changeState int
	for i, seat := range allSeats {
		newSeats[i] = seat
		s := newSeats[i]

		if seat.empty {
			for _, neighboar := range seat.neighboars {
				if allSeats[neighboar].empty {
					s.empty = false
				}
			}
		} else {
			changeState = 0
			for _, neighboar := range seat.neighboars {
				if !allSeats[neighboar].empty {
					changeState++
				}
			}
			if changeState > 3 {
				s.empty = true
			}
		}

		newSeats[i] = s
	}
	return newSeats
}

func main() {
	dat, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var allSeats seats
	allSeats = make(seats)
	var line string
	row := 0
	for scanner.Scan() {
		line = scanner.Text()
		allSeats.getSeats(line, row)
		row++
	}
	col := len(line)

	neighborSeats := allSeats.getNeighbors(row, col)
	allSeats.printSeats(row, col)
	fmt.Printf("%+v\n", neighborSeats)
	//fmt.Printf("%+v\n", allSeats.getNeighbors(3, 3))

}
