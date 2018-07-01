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

	data, err := readDataFromCsv(csvPath)
	if err != nil {
		return
	}

	answers := getAnswers(data, getUserInput)

	fmt.Printf(answers[0])

}

func readDataFromCsv(path string) ([][]string, error) {
	if path[len(path)-4:] != ".csv" {
		errMsg := "Error: Questions and answers must be in a .csv file, received " + path[len(path)-4:]
		err:= errors.New(errMsg)
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
	data, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Printf("successfully read %d lines from %s \n", len(data), path)
	return data, nil
}

 func getAnswers(data [][]string, getInput inputGetter) []string {
	
  answers := make([]string, len(data))

  for i, v := range data{
		fmt.Printf("Question %d of %d:\n", i+1, len(data))
		fmt.Println(v[0])
		a := getInput()
		fmt.Println(a)
		answers[i] = a
	}
  return answers
 }


 func getUserInput() string{
	fmt.Printf("answer:" )
	text := ""
	fmt.Scanln(&text)
	return text
	}



//func markQuiz(data [][]string, answers)(totalq, score,)
	



