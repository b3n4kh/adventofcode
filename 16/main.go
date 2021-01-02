package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	name          string
	lowerRangeMin int
	lowerRangeMax int
	upperRangeMin int
	upperRangeMax int
}

func getTicket(line string) (tickets []int) {
	ticketLine := strings.Split(line, ",")
	for _, t := range ticketLine {
		ticket, _ := strconv.Atoi(t)
		tickets = append(tickets, ticket)
	}
	return tickets
}

func main() {
	dat, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var line string
	var myTickets, nearbyTickets []int
	for scanner.Scan() {
		line = scanner.Text()
		if line == "your ticket:" {
			scanner.Scan()
			myTickets = getTicket(scanner.Text())
		}
		if line == "nearby tickets:" {
			scanner.Scan()
			nearbyTickets = getTicket(scanner.Text())
		}

		fmt.Printf("my ticket: %+v\n", myTickets)
		fmt.Printf("nearby ticket: %+v\n", nearbyTickets)
	}
}
