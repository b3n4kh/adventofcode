package main

import "testing"

func Test_main(t *testing.T) {
	seat := parseBoardingPass("FBFBBFFRLR")

	if seat.row != byte(44) {
		t.Errorf("wrong row = %v, want %v", seat.row, 44)
	}

	if seat.col != byte(5) {
		t.Errorf("wrong col = %v, want %v", seat.col, 5)
	}

	if seat.getID() != 357 {
		t.Errorf("wrong seatID = %v, want %v", seat.getID(), 357)
	}
}
