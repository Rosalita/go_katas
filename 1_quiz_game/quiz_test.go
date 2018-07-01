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

func TestReadDataFromCsvErrors(t *testing.T) {
	var tests = []struct {
		input string
		err   error
	}{
		{"testdata/test.csv", nil},
		{"testdata/helloworld.txt", errors.New("Error: Questions and answers must be in a .csv file, received .txt")},
	}
	for _, test := range tests {
		_, e := readDataFromCsv(test.input)
		assert.Equal(t, e, test.err, "unexpected error returned")
	}
}

func TestReadDataFromCsvReturnsDataOfCorrectLength(t *testing.T) {
	var tests = []struct {
		input      string
		dataLength int
	}{
		{"testdata/test.csv", 2},
		{"testdata/missing.csv", 0},
		{"testdata/helloworld.txt", 0},
	}
	for _, test := range tests {
		data, _ := readDataFromCsv(test.input)
		assert.Equal(t, len(data), test.dataLength, "unexpected data returned")
	}
}

func TestCanReadDataFromValidCsv(t *testing.T) {
	path := "testdata/test.csv"
	data, _ := readDataFromCsv(path)

	expected_row1 := []string{"1 + 1", "2"}
	expected_row2 := []string{"Why can't spaghetti code?", "Impasta syndrome"}
	expected := [][]string{expected_row1, expected_row2}

	for j, row := range data {
		for i, item := range row {
			assert.Equal(t, expected[j][i], item, "unexpected values read from CSV")
		}
	}
}

func TestCsvFileIsMissing(t *testing.T) {
	path := "testdata/missing.csv"
	_, err := readDataFromCsv(path)

	expected := "open testdata/missing.csv: no such file or directory"
	assert.Equal(t, expected, err.Error(), "unexpected error message")
}
