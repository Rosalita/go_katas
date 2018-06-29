package main

import "fmt"
import "flag"

// import "encoding/csv"
import "os"

// global variable for csv flag
var csv string

func main() {

	
	flag.StringVar(&csv, "csv", "problems.csv", "the name of the csv to use")
	flag.Parse()

	fmt.Println(csv)

	csv := getCsvName(os.Args)
	fmt.Println(csv)

	// open the csv file by default
	// create a new reader
	// pass this reader to ther new csv reader method
	// use read() and test EOF

}

func getCsvName(args []string) string {
	if len(args) == 2 {
		return args[1]
	} else {
		return "problems.csv"
	}
}
