package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
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

func parseRule(line string) (name string, parsedRule rule) {
	sub := strings.Split(line, ": ")
	name = sub[0]
	rules := strings.Split(sub[1], " or ")
	lowerRange := strings.Split(rules[0], "-")
	upperRange := strings.Split(rules[1], "-")

	lowerRangeMin, _ := strconv.Atoi(lowerRange[0])
	lowerRangeMax, _ := strconv.Atoi(lowerRange[1])
	upperRangeMin, _ := strconv.Atoi(upperRange[0])
	upperRangeMax, _ := strconv.Atoi(upperRange[1])

	parsedRule = rule{
		lowerRangeMin,
		lowerRangeMax,
		upperRangeMin,
		upperRangeMax,
	}
	return name, parsedRule
}

func checkRule(r rule, value int) bool {
	if (value >= r.lowerRangeMin && value <= r.lowerRangeMax) ||
		(value >= r.upperRangeMin && value <= r.upperRangeMax) {
		return true
	}
	return false
}

func isInVaildTicket(rules map[string]rule, ticket int) bool {
	for _, r := range rules {
		if checkRule(r, ticket) {
			return false
		}
	}
	return true
}

func valuePassesRule(value int, ruleBounds [2][2]int) bool {
	firstBounds := ruleBounds[0]
	secondBounds := ruleBounds[1]
	return ((value >= firstBounds[0] && value <= firstBounds[1]) ||
		(value >= secondBounds[0] && value <= secondBounds[1]))
}

func findMatchingRules(tickets [][]int, rules map[string]rule) (matchingRules map[string]int) {
	skipTicket := map[int]bool{}
	matchingRules = map[string]int{}
	for len(rules) > 0 {

		for ticketIndex := range tickets[0] {
			if skipTicket[ticketIndex] {
				continue
			}
			var passingNames []string
			for ruleName, r := range rules {
				allValuesPassed := true
				// iterate over all tickets and if any fail for this rule, break out
				for _, ticket := range tickets {
					if !(checkRule(r, ticket[ticketIndex])) {
						allValuesPassed = false
						break
					}
				}
				if allValuesPassed {
					passingNames = append(passingNames, ruleName)
				}
			}
			if len(passingNames) == 1 {
				matchingRules[passingNames[0]] = ticketIndex
				// remove the rule from the map b/c we've determined its index
				delete(rules, passingNames[0])
				// remember which indices have already been taken by a rule
				skipTicket[ticketIndex] = true
			}
		}
	}
	return matchingRules
}

func findInvaildTickets(tickets [][]int, rules map[string]rule) (invalidColumns []int, validTickets [][]int) {
	for _, ticket := range tickets {
		isvalid := true
		for _, column := range ticket {
			isinvalid := isInVaildTicket(rules, column)
			if isinvalid {
				invalidColumns = append(invalidColumns, column)
				isvalid = false
				continue
			}
		}
		if isvalid {
			validTickets = append(validTickets, ticket)
		}
	}
	return invalidColumns, validTickets
}

func main() {
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	nbt := false
	inRuleBlock := true
	var line string
	rules := make(map[string]rule)
	var myTickets []int
	var nearbyTickets [][]int
	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			inRuleBlock = false
		}
		if inRuleBlock {
			name, rule := parseRule(line)
			rules[name] = rule
		}
		if line == "your ticket:" {
			scanner.Scan()
			myTickets = getTicket(scanner.Text())
		}
		if line == "nearby tickets:" {
			nbt = true
			continue
		}
		if nbt {
			nearbyTickets = append(nearbyTickets, getTicket(line))
		}
	}
	invalidTickets, validTickets := findInvaildTickets(nearbyTickets, rules)
	//fmt.Printf("validTickets: %+v\n", validTickets)
	//fmt.Printf("invalidTickets: %+v\n", invalidTickets)

	fieldNameToIndex := findMatchingRules(validTickets, rules)
	fmt.Printf("fieldNameToIndex: %+v\n", fieldNameToIndex)
	result1 := 0
	result2 := 1
	for _, ticket := range invalidTickets {
		result1 += ticket
	}
	for name, index := range fieldNameToIndex {
		if strings.HasPrefix(name, "departure") {
			result2 *= myTickets[index]
		}
	}
	fmt.Printf("Result 1: %v\n", result1)
	fmt.Printf("Result 2: %+v\n", result2)

	//fmt.Printf("invalid tickets: %+v\n", invalidTickets)
}
