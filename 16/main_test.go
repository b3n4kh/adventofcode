package main

import (
	"fmt"
	"strings"
	"testing"
)


func Test_main(t *testing.T) {
	dat, err := os.Open("test.txt")
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
