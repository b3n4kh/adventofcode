package main

import (
	"bufio"
	"fmt"
	"os"
)

type rule struct {
	name          string
	lowerRangeMin int
	lowerRangeMax int
	upperRangeMin int
	upperRangeMax int
}

func main() {
	dat, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if line != "" {

		}
		fmt.Println(line)
	}
}
