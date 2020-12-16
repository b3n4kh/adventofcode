package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func addOutletAndDevice(numbers []int) []int {
	device := numbers[len(numbers)-1] + 3
	numbers = append(numbers, 0)
	numbers = append(numbers, device)
	sort.Ints(numbers)
	return numbers
}

func getdifferences(numbers []int) (oneJolt int, threeJolt int) {
	sort.Ints(numbers)
	numbers = addOutletAndDevice(numbers)
	var difference int
	for i := 0; i < len(numbers)-1; i++ {
		difference = numbers[i+1] - numbers[i]
		//fmt.Printf("%v - %v = %v\n", numbers[i+1], numbers[i], difference)
		switch difference {
		case 1:
			oneJolt++
		case 3:
			threeJolt++
		}
	}
	return oneJolt, threeJolt
}

func getArrangements(numbers []int) (arrangements int) {
	sort.Ints(numbers)
	numbers = addOutletAndDevice(numbers)
	var lastMax int
	arrangedNumbers := make([][]int, len(numbers)-1)
	for i := 0; i < len(numbers)-1; i++ {
		arrangedNumbers[i] = make([]int, len(numbers)-1)
		for diff := 1; diff <= 5; diff++ {
			if len(numbers) > i+diff {
				if i > 0 {
					lastMax = arrangedNumbers[i-1][2] + 3
				} else {
					lastMax = numbers[i] + 3
				}
				//fmt.Printf("numbers[%v + %v] = %+v\t lastmax: %v\n", i, diff, numbers[i+diff], lastMax)

				if numbers[i+diff] <= (lastMax) {
					arrangedNumbers[i][diff-1] = numbers[i+diff]
				}
			}
			if arrangedNumbers[i][diff-1] == 0 {
				arrangedNumbers[i][diff-1] = numbers[i+1]
			}
		}
	}
	fmt.Printf("%+v \n", arrangedNumbers)
	for diff := 0; diff < 3; diff++ {
		for i := range arrangedNumbers {
			element := arrangedNumbers[i][diff]
			if element != 0 {
				fmt.Printf("%+v, ", element)
			}
		}
		fmt.Printf("\n")
	}
	return arrangements
}

func main() {
	dat, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var currNumber int
	var numbers []int

	var line string
	for scanner.Scan() {
		line = scanner.Text()
		currNumber, _ = strconv.Atoi(line)
		numbers = append(numbers, currNumber)
	}
	arrangements := getArrangements(numbers)
	fmt.Printf("\nResult2: %v\n\n", arrangements)
	return

	oneJolt, threeJolt := getdifferences(numbers)
	fmt.Printf("\n1-jolt: %v, 3-jolt: %v\n\n", oneJolt, threeJolt)
	fmt.Printf("Result1: %v", oneJolt*threeJolt)
}
