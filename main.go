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
	flag.StringVar(&filename, "f", "problems.csv", "filename of CSV problem set")
	flag.IntVar(&limit, "limit", 30, "time limit in seconds")
}

func main() {
	flag.Parse()
	err := quiz()
	if err != nil {
		log.Fatalf("Error starting quiz game: %s", err)
	}
}
