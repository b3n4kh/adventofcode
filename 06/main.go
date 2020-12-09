package main

import (
	"bufio"
	"fmt"
	"os"
)

// Questions parsed
type Questions struct {
	votes        map[rune]int
	groupMembers int
}

// SumOfAnswers for output
type SumOfAnswers struct {
	questions int
	inUnison  int
}

func (questions *Questions) parseAnswers(line string) {
	if questions.votes == nil {
		questions.votes = make(map[rune]int)
	}

	for _, char := range line {
		questions.votes[char]++
	}
	questions.groupMembers++
}

func (answers *SumOfAnswers) parseGroups(questions Questions) {
	answers.questions += len(questions.votes)
	for _, i := range questions.votes {
		if i == questions.groupMembers {
			answers.inUnison++
		}
	}
}

func main() {
	dat, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var line string
	answers := SumOfAnswers{}
	questions := Questions{}
	for scanner.Scan() {
		line = scanner.Text()
		if line != "" {
			questions.parseAnswers(line)
			continue
		}
		answers.parseGroups(questions)

		questions = Questions{}
	}

	fmt.Printf("Sum of Questions: %+v\n", answers.questions)
	fmt.Printf("Sum of in Unison Questions: %+v", answers.inUnison)

}
