package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type problem struct {
	q string
	a string
}

func main() {
	// get the file to read from
	fileName := flag.String("file", "problems.csv", "a string of the csv file name with the format of 'question,answer'")

	// get the time limit
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds")

	flag.Parse()
	fmt.Println(*fileName)

	// open the csv file
	csvFile, err := os.Open(*fileName)
	if err != nil {
		exit(fmt.Sprintf("Could not open the file %q \n", *fileName))
	}

	fmt.Println("Successfully opened the file")
	defer csvFile.Close()

	// read the data in the csv file
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		exit("Could not read the file")
	}

	problems := parseQuestions(csvLines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	var score int

problemloop:
	for i, prob := range problems {
		// get the question and answer
		question := prob.q
		answer := strings.TrimSpace(prob.a)

		// print out the questions
		fmt.Printf("question %v: %v \n", i+1, question)

		answerCh := make(chan string)

		go func() {
			var user_input string
			// receive the user's input
			fmt.Scanln(&user_input)
			answerCh <- user_input
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case user_input := <-answerCh:
			if user_input == answer {
				score++
			}
		}
	}

	fmt.Printf("You scored %v out of %v", score, len(problems))
}

func parseQuestions(questions [][]string) []problem {
	prob := make([]problem, len(questions))

	for i, line := range questions {
		prob[i] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return prob
}
