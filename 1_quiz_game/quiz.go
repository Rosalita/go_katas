package main

import "flag"
import "encoding/csv"
import "os"
import "log"
import "errors"
import "fmt"

var csvPath string = ""

type inputGetter func() string

func main() {

	flag.StringVar(&csvPath, "csv", "problems.csv", "path to csv file")
	flag.Parse()

	quizData, err := readFromCsv(csvPath)
	if err != nil {
		return
	}

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

	for i, v := range quizData {
		fmt.Printf("Question %d of %d:\n", i+1, len(quizData))
		fmt.Println(v[0])
		a := getInput()
		answers[i] = a
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
		if quizData[i][1] == answers[i] {
			counter++
		}
	}
	return counter
}
