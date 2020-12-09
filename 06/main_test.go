package main

import (
	"strings"
	"testing"
)

const testdata = `abc

a
b
c

ab
ac

a
a
a
a

b

`

func Test_main(t *testing.T) {

	answers := SumOfAnswers{}
	questions := Questions{}
	for _, line := range strings.Split(testdata, "\n") {
		if line != "" {
			questions.parseAnswers(line)
			continue
		}
		answers.parseGroups(questions)

		questions = Questions{}
	}

	if answers.questions != 11 {
		t.Errorf("Wrong number of questions: %v wanted %v", answers.questions, 11)
	}

	if answers.inUnison != 6 {
		t.Errorf("Wrong number of questions inUnison: %v wanted %v", answers.inUnison, 6)
	}
}
