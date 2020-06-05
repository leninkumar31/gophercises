package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// Read command line flags
	filePath := flag.String("path", "problems.csv", "a string")
	duration := flag.Int("wait", 30, "a string")
	flag.Parse()
	// Open file for reading
	file, err := os.Open(*filePath)
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
		input := make(chan string, 1)
		go func() {
			var msg string
			fmt.Scan(&msg)
			input <- msg
		}()
		exit := make(chan bool, 1)
		go func(duration *int) {
			time.Sleep(time.Second * time.Duration(*duration))
			exit <- true
		}(duration)
		select {
		case answer := <-input:
			if answer == record[1] {
				correctAnswers++
			}
		case <-exit:
			fmt.Println("User didn't give any answer. Proeeding to next question")
		}
	}
	fmt.Printf("You got %v out of %v questions correct\n", correctAnswers, numQuestions)
}
