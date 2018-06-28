package main

import "fmt"

// import "encoding/csv"
import "os"

func main() {

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
