package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.Parse()

	inputfile := "input.txt"

	if debug {
		inputfile = "test.txt"
	}

	dat, err := ioutil.ReadFile(inputfile)
	check(err)
	scanner := bufio.NewScanner(bytes.NewReader(dat))
	current_depth, last_depth, second_last_depth := 0, 0, 0
	current_window, last_window := 0, 0
	result_one, result_two := -1, -1

	for scanner.Scan() {
		current_depth, err = strconv.Atoi(scanner.Text())
		check(err)
		if current_depth > last_depth {
			result_one++
		}

		if second_last_depth > 0 && last_depth > 0 {
			current_window = current_depth + last_depth + second_last_depth
			if current_window > last_window {
				result_two++
			}
			last_window = current_window
		}
		second_last_depth = last_depth
		last_depth = current_depth
	}

	fmt.Printf("Part One: %v\n", result_one)
	fmt.Printf("Part Two: %v\n", result_two)
}
