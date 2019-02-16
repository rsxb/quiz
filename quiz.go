package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func quiz() error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("quiz: failed to open csv file: %s", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return fmt.Errorf("quiz: failed to parse csv file: %s", err)
	}
	n := len(records)

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
		for i, v := range records {
			// question and answer pair
			q, a := v[0], v[1]
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
