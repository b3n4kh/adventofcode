package main

import (
	"strings"
	"testing"
)

const testdata = `F10
N3
F7
R90
F11`

func Test_main(t *testing.T) {
	shipOne := Ship{'E', 0, 0}

	for _, line := range strings.Split(testdata, "\n") {
		shipOne.parseInstructionsOne(line)
	}

	if shipOne.longitudo != -17 || shipOne.latitudo != -8 {
		t.Errorf("The result was: |%v| + |%v| = %v wanted |%v| + |%v| = %v ", shipOne.longitudo, shipOne.latitudo, abs(shipOne.longitudo)+abs(shipOne.latitudo), -17, 8, 25)
	}

}
