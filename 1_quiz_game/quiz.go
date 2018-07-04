package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var csvPath string
var limit int

type inputGetter func() string

func main() {

	flag.IntVar(&limit, "limit", 30, "a time limit for answering the quiz")
	flag.StringVar(&csvPath, "csv", "problems.csv", "path to csv file")
	flag.Parse()

	quizData, err := readFromCsv(csvPath)
	if err != nil {
		return
	}

	fmt.Println("Press <Enter> to start quiz")
	fmt.Scanln()

	answers := getAnswers(quizData, getUserInput)
	score := markQuiz(quizData, answers)
	fmt.Printf("Your score was %d out of %d correct\n", score, len(quizData))
}

func readFromCsv(path string) ([][]string, error) {
	if path[len(path)-4:] != ".csv" {
		errMsg := "Error: Questions and answers must be in a .csv file, received " + path[len(path)-4:]
		err := errors.New(errMsg)
		log.Println(err)
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	quizData, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Printf("successfully read %d lines from %s \n", len(quizData), path)
	return quizData, nil
}

func getAnswers(quizData [][]string, getInput inputGetter) []string {

	answers := make([]string, len(quizData))
	timer := time.NewTimer(time.Duration(limit) * time.Second)

	for i, v := range quizData {
		fmt.Printf("Question %d of %d:\n", i+1, len(quizData))
		fmt.Println(v[0])

		answerCh := make(chan string)
		go func() {
			a := getInput()
			answerCh <- a
		}()
		select {
		case <-timer.C:
			fmt.Println("\n *** Time up ***")
			return answers
		case a := <-answerCh:
			answers[i] = a

		}
	}
	return answers
}

func getUserInput() string {
	fmt.Printf("answer:")
	text := ""
	fmt.Scanln(&text)
	return text
}

func markQuiz(quizData [][]string, answers []string) int {
	counter := 0

	for i, _ := range quizData {
		quizAnswer := strings.ToLower(quizData[i][1])
		userAnswer := strings.ToLower(answers[i])

		if userAnswer == quizAnswer {
			counter++
		}
	}
	return counter
}
