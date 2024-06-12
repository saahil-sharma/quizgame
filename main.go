package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func main() {

	//Uses the imported flag package to add a flag to the command line.
	//Allows users to change the value of the name of the csv file at the command line.
	//Second parameter sets the default value to "problems.csv".
	//flag.Parse() is needed at the end of all flag.* calls.
	csvFilename := flag.String("csv", "problems.csv", "a csv file")
	flag.Parse()

	//Open the file.
	file, err := os.Open(*csvFilename)

	//Check if there was an error with opening the file.
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file: %s\n", *csvFilename))
		os.Exit(1)
	}

	//Reads the csv file and splits the file into a 2D array.
	//Each array within lines contains all the values in the line of the csv file.
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file. ")
	}
	problems := parseLines(lines)

	counter := 0
	for i, p := range problems {
		fmt.Printf("Problem: #%d: %s = ", i+1, p.q)
		var answer string
		_, err2 := fmt.Scanf("%s\n", &answer)
		if err2 != nil {
			return
		}
		if answer == p.a {
			counter++
		}
	}
	fmt.Printf("You scored %d out of %d. ", counter, len(problems))
	fmt.Println(getMood(counter, problems))
}

// Function to parse the lines of a 2D array.
// Takes each array within lines and turns them into a
// struct of type problem, storing the problem and answer.
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// /Function to print the message when there is an error.
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func getMood(answer int, problems []problem) string {
	if answer > len(problems)/2.0 {
		return "Nice job!"
	} else {
		return "Work Harder!"
	}
}
