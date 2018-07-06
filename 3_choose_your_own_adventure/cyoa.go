package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func main() {

	story := generateStory("./story.json")

	//fmt.Println("story struct: %+v", story)
	// fmt.Println(story["new-york"].Title)
	// fmt.Println(story["new-york"].Paragraphs[1])
	fmt.Println(story["new-york"].Options[1].Chapter)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func generateStory(pathToJSON string) Story {
	data, err := ioutil.ReadFile(pathToJSON)
	if err != nil {
		log.Print(err)
		return nil
	}
	var story Story
	json.Unmarshal(data, &story)
	return story
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// Choose your own adventure - TDD
// sample json - done
// read it in - done
// unmarshal it into a struct - done
// json to go = https://mholt.github.io/json-to-go/
// create a http.Handler to handle web requests instead of a handler function
// use html/template package to create html pages that display the story
