package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Password after parsing input
type Password struct {
	minKey   int
	maxKey   int
	letter   string
	password string
}

func parseFields(fields []string) Password {
	keyRange := strings.Split(fields[0], "-")
	minKey, err := strconv.Atoi(keyRange[0])
	maxKey, err := strconv.Atoi(keyRange[1])
	check(err)
	key := string(fields[1][0])
	password := fields[2]

	return Password{minKey, maxKey, key, password}
}

func checkPassword(password Password) int {
	keyCount := strings.Count(password.password, password.letter)
	if keyCount >= password.minKey && keyCount <= password.maxKey {
		return 1
	}

	return 0
}

func checkPassword2(password Password) int {
	if (string(password.password[password.minKey-1]) == password.letter) != (string(password.password[password.maxKey-1]) == password.letter) {
		return 1
	}

	return 0
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)
	var result1, result2 int
	scanner := bufio.NewScanner(bytes.NewReader(dat))
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		password := parseFields(fields)
		result1 = result1 + checkPassword(password)
		result2 = result2 + checkPassword2(password)
	}
	fmt.Printf("Number of Matches for 1: %d\n", result1)
	fmt.Printf("Number of Matches for 2: %d", result2)
}
