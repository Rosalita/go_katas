package main

import "testing"
import "flag"
import "os"
import "errors"

import "github.com/stretchr/testify/assert"

func TestMain(m *testing.M) {

	flag.Parse()
	exitCode := m.Run()
	os.Exit(exitCode)

}

// to do refactor length of responses into a test table

func TestReadDataFromCsvErrors(t *testing.T){
	var tests = []struct{
		input string
		err error 
	}{
		{"testdata/test.csv", nil},
		{"testdata/helloworld.txt", errors.New("Error: Questions and answers must be in a .csv file, received .txt")},
	}
	for _, test := range tests {
		_, e := readDataFromCsv(test.input); 
		assert.Equal (t, e, test.err, "unexpected error returned")
	}
} 

func TestCanReadDataFromValidCsv(t *testing.T) {
	path := "testdata/test.csv"
	data, _ := readDataFromCsv(path)

	if len(data) != 2 {
		t.Errorf("Error was expecting data with length of 2 but length was: %d", len(data))
	}

	expected_row1 := []string{"1 + 1", "2"}
	expected_row2 := []string{"Why can't spaghetti code?", "Impasta syndrome"}
	expected := [][]string{expected_row1, expected_row2}

	for j, row := range data {
		for i, item := range row {
			assert.Equal (t, item, expected[j][i], "unexpected values read from CSV")
		}
	}
}

func TestCsvFileIsMissing(t *testing.T) {
	path := "testdata/missing.csv"
	data, err := readDataFromCsv(path)

	if len(data) != 0 {
		t.Errorf("Error was expecting data with length of 0 but length was: %d", len(data))
	}

	expected := "open testdata/missing.csv: no such file or directory"
	assert.Equal (t, err.Error(), expected, "unexpected error message")
}

func TestQuestionsNotInCsvFormat(t *testing.T) {
	path := "testdata/helloworld.txt"
	data, err := readDataFromCsv(path)

	if len(data) != 0 {
		t.Errorf("Error was expecting data with length of 0 but length was: %d", len(data))
	}

	expected := "Error: Questions and answers must be in a .csv file, received .txt"
	assert.Equal (t, err.Error(), expected, "unexpected error message")
}

