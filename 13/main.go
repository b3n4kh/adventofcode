package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getBusIDs(line string) (busIDs []int) {
	split := strings.Split(line, ",")
	for _, bus := range split {
		if bus != "x" {
			busID, _ := strconv.Atoi(bus)
			busIDs = append(busIDs, busID)
		}
	}
	return busIDs
}

func getNextBus(currentTime int, busIDs []int) (nextBus int, departure int) {
	departure = currentTime
	for {
		for _, bus := range busIDs {
			if (departure % bus) == 0 {
				return bus, departure
			}
		}
		departure++
	}
}

func getFirstTimestampWalk(busIDs []int) (timestamp int) {

	return timestamp
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	contents := string(dat)
	lines := strings.Split(contents, "\n")

	currentTime, _ := strconv.Atoi(lines[0])
	busIDs := getBusIDs(lines[1])
	bus, departure := getNextBus(currentTime, busIDs)
	timoToWait := departure - currentTime

	firstTimestampWalk := getFirstTimestampWalk(busIDs)
	//fmt.Printf("%+v\ntimestamp: %v\n", busIDs, currentTime)
	//fmt.Printf("Next Bus: %+v\tDeparts at: %v\n", bus, departure)
	fmt.Printf("Time to wait: %v\tNext Bus ID: %v\nResult: %v\n\n", timoToWait, bus, timoToWait*bus)
	fmt.Printf("First Timestamp Walk: %v", firstTimestampWalk)

}
