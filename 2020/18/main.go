package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Operand interface
type Operand interface {
	doMath() int
}

// Operation interface
type Operation struct {
	Operands  []Operand
	Operators []Operator
}

// Operator interface
type Operator func(int, int) int

func sum(x, y int) int {
	return x + y
}

func product(x, y int) int {
	return x * y
}

func (o Operation) doMath() (result int) {
	for i := range o.Operands {
		if o.Operators[i](1, 1) == 2 {

		}
	}

	return result
}

func parseParanthesis(elements []string) (subElements map[int]string) {
	var depth, lastDepth, subBlock, diff int
	var numOrOp string
	subElements = make(map[int]string)

	for _, e := range elements {
		for _, char := range e {
			switch char {
			case '(':
				depth++
			case ')':
				depth--
			default:
				numOrOp = string(char)
			}
		}
		//fmt.Printf("depth: %v, lastDepth: %v, numOrOp: %v\n", depth, lastDepth, numOrOp)
		diff = depth - lastDepth
		lastDepth = depth

		if diff > 0 {
			subBlock += diff
			subElements[subBlock] += numOrOp
		} else {
			subElements[subBlock] += numOrOp
			subBlock += diff
		}
	}
	return subElements
}

func calculate(line string) (result int) {
	line = strings.ReplaceAll(line, "(", "( ")
	line = strings.ReplaceAll(line, ")", " )")

	elements := strings.Split(line, " ")
	//subElements := parseParanthesis(elements)
	var numberL, numberR, depth int
	var operator rune
	for i, e := range elements {

		for _, char := range e {
			switch char {
			case '(':
				depth++
			case ')':
				depth--
			case '*', '+':
				operator = char
			default:
				if numberL == 0 {
					numberL, _ = strconv.Atoi(string(char))
					continue
				}
				numberR, _ = strconv.Atoi(string(char))
			}
		}
		if numberL != 0 && numberR != 0 {
			if operator == '+' {
				result += numberL + numberR
			}
			if operator == '*' {
				result += numberL * numberR
			}
		}
		fmt.Printf("Index: %v, Element: %v\n", i, e)

	}
	return result
}

func main() {
	dat, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var line string
	var result int
	for scanner.Scan() {
		line = scanner.Text()
		result += calculate(line)
	}
	fmt.Printf("Result1: %v\n", result)
}
