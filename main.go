package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file")
	flag.Parse()
	_ = csvFilename

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Failed to open the csv file: %s: " * csvFilename)
		os.Exit(1)
	}
	_ = file
}
