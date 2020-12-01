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

func find2020(data []int) {
	for first, firstNum := range data {
		for second, secondNum := range data[first:] {
			for _, thirdNum := range data[first+second:] {
				if (firstNum + secondNum + thirdNum) == 2020 {
					result := firstNum * secondNum * thirdNum
					fmt.Printf("%v * %v * %v = %v\n", firstNum, secondNum, thirdNum, result)
				}
			}
		}
	}
}

func main() {

	dat, err := ioutil.ReadFile("01.txt")
	check(err)
	var result []int
	scanner := bufio.NewScanner(bytes.NewReader(dat))
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		check(err)
		result = append(result, x)
	}
	find2020(result)
}
