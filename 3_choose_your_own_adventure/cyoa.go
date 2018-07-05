package main

import(
"fmt"
//"errors"
//"io/ioutil"
"encoding/json"
)

type StoryOptions struct {
	Text string `json:"text"`
	Arc string `json:"arc"`
}

type Story struct {
	Title string `json:"title"`
	Options []StoryOptions
}

// type Story struct {
// 	Title string `json:"title"`
// 	StoryText []string `json:"story"`
// 	Options []StoryOptions
// }

type Bird struct {
	Species string
	Description string
}

func main(){

	somejson := `{"species": "pigeon","description": "likes to perch on rocks"}`

	var bird Bird	
    json.Unmarshal([]byte(somejson), &bird)
	fmt.Printf("Species: %s, Description: %s\n", bird.Species, bird.Description)

	morejson := `[{"text": "Let's head to New York.","arc": "new-york"},{"text": "Let's try our luck in Denver.","arc": "denver"}]`

	var storyoptions []StoryOptions
	json.Unmarshal([]byte(morejson), &storyoptions)

	//fmt.Printf("text: %s arc: %s \n", storyoptions.Text, storyoptions.Arc)
	fmt.Printf("stuff: %+v \n", storyoptions)

	storyjson:= `{"title": "Test Story","options":[{"text": "Let's head to New York.","arc": "new-york"},{"text": "Let's try our luck in Denver.","arc": "denver"}]`
	
	var story Story
	json.Unmarshal([]byte(storyjson), &story)
	fmt.Printf("stuff: %+v \n", story)


	sliceOfStrings := []string{"apple", "peach", "pear"}
	json, _ := json.Marshal(sliceOfStrings)
	fmt.Println(string(json))

	// temp, _ := ioutil.ReadAll(story.json)

	// fmt.Println(temp)

	//var adventure []Story



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

