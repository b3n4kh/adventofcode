package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// BagRule containment
type BagRule struct {
	color    string
	contains map[string]int
}

func parseRule(line string) (rule BagRule) {
	lineSlice := strings.Split(line[:len(line)-1], " bags contain ")
	containRules := strings.Split(lineSlice[1], ", ")
	contains := make(map[string]int)

	for _, rule := range containRules {
		tokens := strings.Split(rule, " ")
		color := tokens[1] + " " + tokens[2]
		numBags, err := strconv.Atoi(tokens[0])
		if err != nil {
			if tokens[0] == "no" {
				numBags = 0
			}
		}
		contains[color] = numBags
	}

	//fmt.Printf("Rules: %+v\n", BagRule{lineSlice[0], contains})

	return BagRule{lineSlice[0], contains}
}

func findContainer(color string, rules []BagRule, containers map[string]int) int {
	if containers == nil {
		containers = make(map[string]int, len(rules))
	}
	for _, rule := range rules {
		for containedColor := range rule.contains {
			if containedColor == color {
				fmt.Printf("Color: %v found in: %+v\n", color, rule)
				containers[rule.color] = 1
				findContainer(rule.color, rules, containers)
			}
		}
	}
	return len(containers)
}

func main() {
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var line string
	var rules []BagRule
	for scanner.Scan() {
		line = scanner.Text()
		rules = append(rules, parseRule(line))
	}

	fmt.Println(findContainer("shiny gold", rules, nil))
}
