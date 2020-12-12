package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

// vibrant plum    = 2x 11 =>  5 faded blue + 6 dotted black
// dark olive bags =    7  =>  3 faded blue + 4 dotted black
// shiny gold      =    3  =>  1 dark olive + 2 vibrant plum

// dark olive bags contain 3 faded blue bags, 4 dotted black bags.

//darkolive     darkolive.contains    vibrantplum    vibrantplum.contains          result
// `   1 +             1*7       +      2      +          2*11                     = 32`

func findContent(color string, rules map[string]map[string]int) int {
	var newContainers, nestedContainers, sumContainers int
	for containerColor := range rules[color] {

		nestedContainers = findContent(containerColor, rules)
		newContainers += rules[color][containerColor]

		if nestedContainers != 0 && newContainers != 0 {
			sumContainers += rules[color][containerColor]*nestedContainers + rules[color][containerColor]
		}
		if nestedContainers == 0 && newContainers != 0 {
			sumContainers += rules[color][containerColor]
		}

		// fmt.Printf("%v => %v =>  Add %v + %v = %v\n", color, containedColor, containedItems, contains, newContainers)
		// result += rules[color][containerColor]
		// fmt.Printf("%v => %v => Contains: %+v\n", color, containerColor, anzahlAnContainern)
	}
	if sumContainers == 0 {
		sumContainers = sumContainers + newContainers
	} else {
		fmt.Printf("Color: %+v => sumContainers %+v\n\n", color, sumContainers)
	}

	return sumContainers
}

func main() {
	dat, err := os.Open("input.txt")
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
	containerblubb := findContent("shiny gold", rules)
	fmt.Printf("Part Two: %+v\n", containerblubb)

	//	fmt.Printf("Part Two: %+v\n", sumContent(containerblubb))
}
