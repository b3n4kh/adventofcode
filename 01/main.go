package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func find2020(data []int) (two int, three int) {
	for first, firstNum := range data {
		for second, secondNum := range data[first:] {
			if (firstNum + secondNum) == 2020 {
				two = firstNum * secondNum
			}
			for _, thirdNum := range data[first+second:] {
				if (firstNum + secondNum + thirdNum) == 2020 {
					three = firstNum * secondNum * thirdNum
					// fmt.Printf("%v * %v * %v = %v\n", firstNum, secondNum, thirdNum, result)
					return two, three
				}
			}
		}
	}
	return two, three
}

func main() {

	dat, err := ioutil.ReadFile("input.txt")
	check(err)
	var result []int
	scanner := bufio.NewScanner(bytes.NewReader(dat))
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		check(err)
		result = append(result, x)
	}
	two, three := find2020(result)

	fmt.Printf("Part One: %v\n", two)
	fmt.Printf("Part Two: %v\n", three)

}
