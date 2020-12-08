package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dat, err := os.Open("input.txt")
	check(err)
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		fmt.Println(line)
	}
}
