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

func findContainerLegacy(color string, rules []BagRule, containers map[string]int) int {
	if containers == nil {
		containers = make(map[string]int, len(rules))
	}
	for _, rule := range rules {
		for containedColor := range rule.contains {
			if containedColor == color {
				// fmt.Printf("Color: %v found in: %+v\n", color, rule)
				containers[rule.color] = 1
				findContainerLegacy(rule.color, rules, containers)
			}
		}
	}
	return len(containers)
}

func parseRule(line string) (color string, rule map[string]int) {
	lineSlice := strings.Split(line[:len(line)-1], " bags contain ")
	containRules := strings.Split(lineSlice[1], ", ")
	contains := make(map[string]int)

	for _, rule := range containRules {
		tokens := strings.Split(rule, " ")
		color := tokens[1] + " " + tokens[2]
		numBags, err := strconv.Atoi(tokens[0])
		if err != nil {
			if tokens[0] == "no" {
				continue
			}
		}
		contains[color] = numBags
	}

	return lineSlice[0], contains
}

func findContainer(color string, rules map[string]map[string]int, containers map[string]int) int {
	if containers == nil {
		containers = make(map[string]int, len(rules))
	}
	for containerColor := range rules[color] {
		// fmt.Printf("Color: %+v found in: %+v\n", containerColor, color)
		containers[containerColor] = 1
		findContainer(containerColor, rules, containers)
	}
	return len(containers)
}

func findContent(color string, rules map[string]map[string]int, containers int) int {
	var contains int
	for containerColor := range rules[color] {
		contains = findContent(containerColor, rules, containers)
		//fmt.Printf("%v => %v =>  Add %v + %v\n", color, containedColor, containedItems, contains)
		//newContainers := containedItems + contains
		//fmt.Printf("%v => %v =>  Add %v + %v = %v\n", color, containedColor, containedItems, contains, newContainers)

		containers += contains
	}
	return containers
}

func main() {
	dat, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var line string
	rules := make(map[string]map[string]int)
	var color string
	rule := make(map[string]int)
	for scanner.Scan() {
		line = scanner.Text()
		color, rule = parseRule(line)
		rules[color] = rule
	}

	fmt.Printf("Part One: %v\n", findContainer("shiny gold", rules, nil))
	// fmt.Printf("Part Two: %v\n", findContent("shiny gold", rules, 1))
}
