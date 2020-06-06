package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Read command line flags
	filePath := flag.String("filepath", "problems.csv", "a csv file in the format of question,answer")
	duration := flag.Int("waitTime", 30, "a string")
	flag.Parse()
	// Open file for reading
	file, err := os.Open(*filePath)
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file: %s\n", *filePath))
	}
	defer file.Close()
	// Read the contents of CSV file
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided csv file")
	}
	problems := parseLines(lines)
	// Print the questoin and Take user input
	correct := 0
	timer := time.NewTimer(time.Duration(*duration) * time.Second)
loop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s=", i+1, p.question)
		input := make(chan string)
		go func() {
			var msg string
			fmt.Scan(&msg)
			input <- msg
		}()
		select {
		case answer := <-input:
			if answer == p.answer {
				correct++
			}
		case <-timer.C:
			break loop
		}
	}
	fmt.Printf("\n%d out of %d are correct\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
