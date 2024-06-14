package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
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
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	counter := 0

problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem: #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			if answer == p.a {
				counter++
			}
		}
	}

	fmt.Printf("You scored %d out of %d. %s", counter, len(problems), getMood(counter, problems))
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
