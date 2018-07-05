package main

import(
"fmt"
//"errors"
"io/ioutil"
"encoding/json"
)

type StoryOptions struct {
	Text string `json:"text"`
	Arc string `json:"arc"`
}

type Story struct {
	Title string
	Story []string
	Options []StoryOptions
}


func main(){
	sliceOfStrings := []string{"apple", "peach", "pear"}
	json, _ := json.Marshal(sliceOfStrings)
	fmt.Println(string(json))

	temp, _ := ioutil.ReadAll(story.json)

	var adventure []Story


	// data, err := readStory("story.json")

	// decoding JSON data into Go values
	// fmt.Println(data)
	// fmt.Println(err)

// Choose your own adventure - TDD
// sample json
// read it in
// marshal it into a struct
// use html/template package to create html pages
// create a http.Handler to handle web requests instead of a handler function
// use encoding/json package to decode the JSON file
// all stories will have a story arc named "intro"
// json to go = https://mholt.github.io/json-to-go/

}

// func readStory(path string)(data string, err error){

// return "hi", errors.New("ow hai") 

