package main

import (
	"flag"
	"log"
)

var (
	filename string
	limit    int
)

func init() {
	flag.StringVar(&filename, "f", "problems.csv", "a csv file in 'question,answer' format")
	flag.IntVar(&limit, "limit", 30, "time limit in seconds")
}

func main() {
	flag.Parse()
	err := Quiz()
	if err != nil {
		log.Fatalf("[ERROR] %s", err)
	}
}
