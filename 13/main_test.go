package main

import (
	"strconv"
	"strings"
	"testing"
)

const testdata = `939
7,13,x,x,59,x,31,19`

func Test_main(t *testing.T) {
	lines := strings.Split(testdata, "\n")
	currentTime, _ := strconv.Atoi(lines[0])
	busIDs := getBusIDs(lines[1])
	bus, departure := getNextBus(currentTime, busIDs)
	timoToWait := departure - currentTime

	if (timoToWait * bus) != 295 {
		t.Errorf("The result was: %v wanted %v ", (timoToWait * bus), 295)
	}
}
