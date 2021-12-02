package main

import (
	"strings"
	"testing"
)

const testdata = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

const testdata2 = `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`

func Test_main(t *testing.T) {
	rules := make(map[string]map[string]int)
	var color string
	rule := make(map[string]int)

	for _, line := range strings.Split(testdata, "\n") {
		color, rule = parseRule(line)
		rules[color] = rule
	}
	containers := (findContainer("shiny gold", rules, nil))

	if containers != 0 {
		t.Errorf("Wrong number of containers: %v wanted %v", containers, 0)
	}
}

func Test_Nest(t *testing.T) {
	rules := make(map[string]map[string]int)
	var color string
	rule := make(map[string]int)

	for _, line := range strings.Split(testdata2, "\n") {
		color, rule = parseRule(line)
		rules[color] = rule
	}
	containers := (findContainer("shiny gold", rules, nil))

	if containers != 0 {
		t.Errorf("Wrong number of containers: %v wanted %v", containers, 0)
	}
}
