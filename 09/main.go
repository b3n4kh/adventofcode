package main

import (
	"bufio"
	"fmt"
	"os"
)

type preamble struct {
	numbers []int
}

func main() {
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		fmt.Println(line)
	}
}
