package main

import "testing"
import "flag"
import "os"

func TestMain(m *testing.M) {

	flag.Parse()
	exitCode := m.Run()
	os.Exit(exitCode)

}

// to do:
// refactor asserts with stretchr/testify
// apply table based testing pattern to csv testsg


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
			if item != expected[j][i] {
				t.Errorf("unexpected result: was expecting: %s, actual was: %s", expected[j][i], item)
			}
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
	if err.Error() != expected {
		t.Errorf("unexpected result: was expecting: %s, actual was: %s", expected, err.Error())
	}
}

func TestQuestionsNotInCsvFormat(t *testing.T) {
	path := "testdata/helloworld.txt"
	data, err := readDataFromCsv(path)

	if len(data) != 0 {
		t.Errorf("Error was expecting data with length of 0 but length was: %d", len(data))
	}

	expected := "Error: Questions and answers must be in a .csv file, received .txt"
	if err.Error() != expected {
		t.Errorf("unexpected result: was expecting: %s, actual was: %s", expected, err.Error())
	}
}

