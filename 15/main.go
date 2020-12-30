package main

import (
	"bufio"
	"fmt"
	"os"
)

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
	var numebr int
	var history []int
	for char := strings.Split(line, ",") {
		number, _ = strconv.Atoi(char)
		history = append(history, number)
	}
	fmt.Println(history)

}
