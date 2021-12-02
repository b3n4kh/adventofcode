package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type preamble struct {
	numbers []int
}

func leftRotation(a []int) []int {
	var newArray []int
	newArray = a[1:]
	a = newArray
	return a
}

func isvalidSum(numbers []int, validator int) (index int) {
	//fmt.Printf("Numbers: %+v\nValidator: %v\tLenght: %v\n\n", numbers, validator, len(numbers))

	for i := 0; i <= len(numbers)-1; i++ {
		for y := i + 1; y <= len(numbers)-1; y++ {
			if (numbers[i] + numbers[y]) == validator {
				return 0
			}
		}
	}
	return validator
}

func checkPartOne(numbers []int, currNumber int) (result int) {
	invalid := isvalidSum(numbers[:len(numbers)-1], currNumber)
	return invalid
}

func checkPartTwo(allNumbers []int, input int) (index int, count int) {
	var sum int

	for {
		if index >= len(allNumbers)-1 {
			return 0, 0
		}
		if allNumbers[index+count] == input {
			count = 0
			sum = 0
			continue
		}
		sum += allNumbers[index+count]
		//fmt.Printf("Sum: %+v\tpostion: %v\t\n\n", sum, index+count)
		count++

		if sum == input {
			return index, count
		}
		if sum > input {
			index++
			// fmt.Printf("Sum: %+v\nindex: %v\tcount: %v\n\n", sum, i, count)

			count = 0
			sum = 0
			continue
		}
	}
}

func main() {
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()

	preamble := 25

	var numbers []int
	var allNumbers []int

	scanner := bufio.NewScanner(dat)
	var currNumber, result, partOne int
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		currNumber, _ = strconv.Atoi(line)
		allNumbers = append(allNumbers, currNumber)

		if len(numbers) <= preamble {
			numbers = append(numbers, currNumber)
			continue
		}

		numbers = leftRotation(numbers)
		numbers = append(numbers, currNumber)

		result = checkPartOne(numbers, currNumber)
		if result != 0 {
			partOne = result
			fmt.Printf("\nResult: %v\n\n", result)
		}
	}
	index, count := checkPartTwo(allNumbers, partOne)
	if index != 0 {
		fmt.Printf("\nResult 2:\t %v + %v\n\n", index, count)
		var s []int
		var sum int
		for i := index; i <= index+count; i++ {
			s = append(s, allNumbers[i])
			sum += allNumbers[i]
			if sum == partOne {
				break
			}
		}
		sort.Ints(s)
		fmt.Printf("Array: %+v", s)
		fmt.Printf("\nResult2: %v + %v = %v\n", s[0], s[len(s)-1], s[0]+s[len(s)-1])
		return
	}

}
