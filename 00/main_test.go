package main

import (
	"fmt"
	"strings"
	"testing"
)

const testdata = `test
test
test`

func Test_main(t *testing.T) {
	for _, line := range strings.Split(testdata, "\n") {
		fmt.Println(line)
	}

}
