package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
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

var story Story

func main() {

	story := generateStory("./story.json")
	fmt.Println(story["intro"].Title)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func generateStory(pathToJSON string) Story {
	data, err := ioutil.ReadFile(pathToJSON)
	if err != nil {
		log.Print(err)
		return nil
	}

	json.Unmarshal(data, &story)
	return story
}

func handler(w http.ResponseWriter, r *http.Request) {

	chapterToRender := r.URL.Path[1:]

	if chapterToRender == "" {
		chapterToRender = "intro"
	}

	t, err := template.ParseFiles("story.html")

	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, story[chapterToRender])

	if err != nil {
		log.Print("template executing error: ", err)
	}
}
