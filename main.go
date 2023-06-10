package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

// For Printing Result
func printResult(no_questions int, score int) {
	fmt.Printf("\nTotal no. of question given: %v \n", no_questions)
	fmt.Printf("Total questions correct: %v \n", score)
	fmt.Printf("You total score is: %v", score)
}

func main() {
	// Flags
	filePath := flag.String("csv", "data.csv", "Path to the CSV file")
	timer_input := flag.Int("timer", 30, "Timer for the Quiz")
	flag.Parse()

	// Checking if the file is given
	if *filePath == "" {
		fmt.Println("Please Provide the file")
		return
	}

	var score int
	var answer string
	var no_questions int

	// Opening and Closing the file
	f, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	// Reading the file line by line
	csvReader := csv.NewReader(f)
	// Setting the timer
	timer := time.NewTimer(time.Duration(*timer_input) * time.Second)

	for {
		data, err := csvReader.Read()
		if err == io.EOF {
			printResult(no_questions, score)
			return
		}
		if err != nil {
			fmt.Println("We had trouble reading the file")
			return
		}
		// Printing the questions
		question := data[0]
		no_questions++
		fmt.Printf("%v ", question)

		answer_chan := make(chan string)

		// Go routine for getting the answer and sending it to the answer channel
		go func() {
			fmt.Scan(&answer)
			answer_chan <- answer

		}()
		select {
		case <-timer.C:
			{
				printResult(no_questions, score)
				return
			}
		case answer := <-answer_chan:
			{
				if answer == data[1] {
					score++
				}

			}
		}

	}

}
