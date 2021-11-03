package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func main() {
	// get the file to read from
	fileName := flag.String("file", "problems.csv", "a string of the csv file name with the format of 'question,answer'")

	flag.Parse()
	fmt.Println(*fileName)

	// open the csv file
	csvFile, err := os.Open(*fileName)
	checkError(err)

	fmt.Println("Successfully opened the file")
	defer csvFile.Close()

	// read the data in the csv file
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	checkError(err)

	var score int

	for i, line := range csvLines {
		// get the question and answer
		question := line[0]
		answer := line[1]

		// print out the questions
		fmt.Printf("question %v: %v \n", i+1, question)
		var user_input string

		// receive the user's input
		fmt.Scanln(&user_input)

		if user_input == answer {
			score++
		}
	}

	fmt.Printf("Your score is %v", score)
}
