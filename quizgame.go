package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	fileName := flag.String("csv", "problems.csv", "A csv file in format of question,answer.")
	d := flag.Int("limit", 30, "The time limit for the quiz in seconds.") //flag variable: duration in seconds
	flag.Parse()
	file, err := os.Open(*fileName)

	if err != nil {
		fmt.Printf("File cannot be opened. Error: %s", err)
		os.Exit(1)
	}
	r := csv.NewReader(file)
	qas, _ := r.ReadAll()
	correctanswers := 0

	go runquiz(qas, &correctanswers)
	time.Sleep(time.Duration(*d) * time.Second)
	fmt.Printf("\n Your score is %d out of %d", correctanswers, len(qas))
}

func runquiz(qas [][]string, correct *int) {
	for i, qa := range qas {
		*correct += quiz(qa, i)
	}
}

func quiz(qa []string, i int) int {
	fmt.Printf("Question %d : %s \n Enter your response: ", i+1, qa[0])
	answer := ""

	_, err := fmt.Scanf("%s \n", &answer)
	if err != nil {
		fmt.Printf("Couldn't accept response. Error: %s. Try again...", err)
		return quiz(qa, i)
	}

	if answer == qa[1] {
		return 1
	} else {
		return 0
	}
}
