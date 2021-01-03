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
	nbt := false
	var line string
	var myTickets, nearbyTickets []int
	for scanner.Scan() {
		line = scanner.Text()

		if line == "your ticket:" {
			scanner.Scan()
			myTickets = getTicket(scanner.Text())
			fmt.Printf("my ticket: %+v\n", myTickets)
		}
		if line == "nearby tickets:" {
			nbt = true
			continue
		}
		if nbt {
			nearbyTickets = getTicket(line)
			fmt.Printf("nearby ticket: %+v\n", nearbyTickets)
		}
	}
}
