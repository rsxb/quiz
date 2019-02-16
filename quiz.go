package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

func convert(record []string) problem {
	return problem{
		question: record[0],
		answer:   record[1],
	}
}

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

	problems := make([]problem, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("parseCSV: %s", err)
		}
		problems = append(problems, convert(record))
	}
	return problems, nil
}

func quiz() error {
	problems, err := parseCSV()
	if err != nil {
		return fmt.Errorf("quiz: %s", err)
	}
	n := len(problems)

	// number of correct answers
	correct := 0
	quit := make(chan int)
	timeout := make(chan bool, 1)

	fmt.Print("Press enter to start...")
	fmt.Scanln()
	go func() {
		time.Sleep(time.Duration(limit) * time.Second)
		timeout <- true
	}()

	go func() {
		for i, v := range problems {
			// question and answer pair
			q, a := v.question, v.answer
			fmt.Printf("Problem #%d: %s = ", i+1, q)

			// user response
			var resp string
			fmt.Scanln(&resp)

			if a == resp {
				correct++
			}
		}
		quit <- 0
	}()

	for {
		select {
		case <-quit:
			fmt.Printf("You scored %d out of %d.", correct, n)
			return nil
		case <-timeout:
			fmt.Printf("\nYou scored %d out of %d.", correct, n)
			return nil
		}
	}
}
