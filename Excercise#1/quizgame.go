package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Open file for reading
	file, err := os.Open("./problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// Read the CSV file until end
	reader := csv.NewReader(file)
	var numQuestions int = 0
	var correctAnswers int = 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		numQuestions++
		fmt.Printf("Quiz %v: %v\n", numQuestions, record[0])
		var answer string
		fmt.Scanln(&answer)
		if answer == record[1] {
			correctAnswers++
		}
	}
	fmt.Printf("You got %v out of %v questions correct\n", correctAnswers, numQuestions)
}
