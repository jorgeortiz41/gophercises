package main

import (
	"encoding/csv" // csv is Go's library for reading and writing CSV files
	"fmt"          // fmt is Go's standard input/output library
	"os"           // os is Go's library for operating system functionality
	"strings"      // strings is Go's library for string manipulation
	"time"         // time is Go's library for time manipulation
)

func getProblems(filename string) map[string]string {
	problemMap := make(map[string]string)

	// Open CSV file and handle error
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()

	// Create new CSV reader
	reader := csv.NewReader(file)

	// Read the CSV
	problems, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Iterate over problems and print each
	for _, problem := range problems {
		question := strings.Trim(problem[0], " ")
		answer := strings.Trim(problem[1], " ")
		problemMap[question] = answer
	}

	return problemMap
}

func quiz(problems map[string]string, seconds int) string { // Return type is a string and an error
	correct := 0
	total := len(problems)
	input := ""
	done := false
	timer := time.NewTimer(time.Duration(seconds) * time.Second)
	stop := make(chan bool)
	go func() {
		<-timer.C
		done = true
		stop <- true

	}()
	for question, answer := range problems {
		fmt.Println(question)
		fmt.Scanln(&input)
		if input == answer {
			correct++
		}
		if done {
			fmt.Println("Time's up!")
			break
		}
	}

	// Print result
	result := fmt.Sprintf("You got %d correct out of %d", correct, total)
	return result
}

func main() {
	// Default values
	filename := "problems.csv"
	seconds := 30

	// Get user input for filename and seconds
	fmt.Println("Welcome to Go-Quiz, Enter a csv file below(.csv) or press enter to use default file(problems.csv):")
	fmt.Scanln(&filename)
	fmt.Println("Enter the duration of quiz(seconds) or press enter to use default(30s):")
	fmt.Scanln(&seconds)

	// Call quiz function and print result
	result := quiz(getProblems(filename), seconds)
	fmt.Println(result)

}
