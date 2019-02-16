package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// problem is a quiz question-answer pair.
type problem struct {
	question string
	answer   string
}

// parseCSV processes csv files and returns quiz problems.
func parseCSV() ([]problem, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("parseCSV: %s", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	if err != nil {
		return nil, fmt.Errorf("parseCSV: %s", err)
	}

	// convert records to problems, line-by-line
	problems := make([]problem, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("parseCSV: %s", err)
		}
		problems = append(
			problems,
			problem{
				question: record[0],
				answer:   record[1],
			},
		)
	}
	return problems, nil
}

// Quiz presents problems to the user.
// Quizzes have configurable time limits.
func Quiz() error {
	problems, err := parseCSV()
	if err != nil {
		return fmt.Errorf("Quiz: %s", err)
	}
	n := len(problems)

	// number of correct answers
	correct := 0
	quit := make(chan int)
	fmt.Print("> Press enter to begin...")
	fmt.Scanln()
	timer := time.NewTimer(time.Duration(limit) * time.Second)

	go func() {
		in := bufio.NewReader(os.Stdin)
		for i, p := range problems {
			q, a := p.question, p.answer
			fmt.Printf("Problem #%d:\t%s = ", i+1, q)

			// read user input
			input, _ := in.ReadString('\n')

			// test for correct answer
			if a == strings.TrimSpace(input) {
				correct++
			}
		}
		quit <- 0
	}()

	for {
		// score quiz if done or >tlimit
		select {
		case <-quit:
			fmt.Printf("> You scored %d out of %d.\n", correct, n)
			return nil
		case <-timer.C:
			fmt.Println("\n> Time's up!")
			fmt.Printf("> You scored %d out of %d.\n", correct, n)
			return nil
		}
	}
}
