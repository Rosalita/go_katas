package main

import "testing"
import "flag"
import "os"

//import "fmt"

func TestMain(m *testing.M) {

	flag.Parse()
	exitCode := m.Run()
	os.Exit(exitCode)

}

func TestDefaultCsvNameCorrect(t *testing.T) {
	args := []string{"./quiz"}
	result := getCsvName(args)
	expected := "problems.csv"
	if result != expected {
		t.Errorf("unexpected result: was expecting: %s, actual was: %s", expected, result)
	}
}

func TestCsvNameCorrectWhenArgPassed(t *testing.T) {
	args := []string{"./quiz", "test.csv"}
	result := getCsvName(args)
	expected := "test.csv"
	if result != expected {
		t.Errorf("unexpected result: was expecting: %s, actual was: %s", expected, result)
	}
}

