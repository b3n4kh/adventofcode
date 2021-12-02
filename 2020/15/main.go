package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const rounds = 30000000

func getStartingNumbers(line string) (startingNumbers []int) {
	var number int
	chars := strings.Split(line, ",")
	for _, char := range chars {
		number, _ = strconv.Atoi(char)
		startingNumbers = append(startingNumbers, number)
	}
	return startingNumbers
}

func turn(history []int) []int {
	turn := len(history)
	lastNum := history[turn-1]
	lastPos := findElement(history, lastNum)

	//fmt.Printf("turn: %v : lastNum: %v next: %v - %v = %v\n", (turn + 1), lastNum, turn, lastPos+1, (turn)-(lastPos+1))

	if lastPos == -1 || lastPos == turn-1 {
		history = append(history, 0)
		return history
	}
	history = append(history, (turn-1)-lastPos)

	return history
}

func findElement(slice []int, element int) (position int) {
	for i := len(slice) - 2; i >= 0; i-- {
		if slice[i] == element {
			return i
		}
	}
	return -1
}

func main() {
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var line string
	scanner.Scan()
	line = scanner.Text()
	history := getStartingNumbers(line)

	for i := len(history); i <= rounds-1; i++ {
		history = turn(history)
	}
	fmt.Printf("Result Part1: %v\n", history[2020])
	fmt.Printf("Result Part2: %v\n", history[len(history)-1])

}
