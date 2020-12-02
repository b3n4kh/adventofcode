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

func parseFields(fields []string) (minKey int, maxKey int, key string, password string) {
	keyRange := strings.Split(fields[0], "-")
	minKey, err := strconv.Atoi(keyRange[0])
	maxKey, err = strconv.Atoi(keyRange[1])
	check(err)
	key = string(fields[1][0])
	password = fields[2]

	return minKey, maxKey, key, password
}

func checkPassword(fields []string) int {
	minKey, maxKey, key, password := parseFields(fields)
	keyCount := strings.Count(password, key)
	if keyCount >= minKey && keyCount <= maxKey {
		return 1
	}

	return 0
}

func checkPassword2(fields []string) int {
	minKey, maxKey, key, password := parseFields(fields)
	if (string(password[minKey-1]) == key) != (string(password[maxKey-1]) == key) {
		return 1
	}

	return 0
}

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)
	var result1 int
	var result2 int
	scanner := bufio.NewScanner(bytes.NewReader(dat))
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		result1 = result1 + checkPassword(fields)
		result2 = result2 + checkPassword2(fields)
	}
	fmt.Printf("Number of Matches for 1: %d\n", result1)
	fmt.Printf("Number of Matches for 2: %d", result2)
}
