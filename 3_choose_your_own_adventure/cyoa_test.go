package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestGenerateStory(t *testing.T) {

	optionsIntro := []Option{Option{Text: "Option 1", Chapter: "Chapter 1"}}
	chapterIntro := Chapter{Title: "Test Intro", Paragraphs: []string{"Para 1", "Para 2"}, Options: optionsIntro}
	options1 := []Option{Option{Text: "Option 2", Chapter: "Chapter 2"}}
	chapter1 := Chapter{Title: "Chapter 1", Paragraphs: []string{"Chapter1 Para1"}, Options: options1}
	expectedStory := make(map[string]Chapter)
	expectedStory["intro"] = chapterIntro
	expectedStory["Chapter 1"] = chapter1

	var tests = []struct {
		path  string
		story Story
	}{
		{"testdata/basicStory.json", expectedStory},
		{"testdata/hello.txt", nil},
	}
	for _, test := range tests {
		data := generateStory(test.path)
		assert.Equal(t, test.story, data, "unexpected data returned")
	}
}

func TestGenerateStoryLogsErrors(t *testing.T) {
	var buffer bytes.Buffer
	log.SetOutput(&buffer)
	generateStory("foo")
	shouldContain := "no such file or directory"
	actual := buffer.String()
	assert.Contains(t, actual, shouldContain, "unexpected log message")
}

func TestCanRenderWebPages(t *testing.T){
	
		request, err := http.NewRequest("GET", "http://localhost:8080/", nil)
		if err != nil {
			t.Fatal(err)
		}
	
		resrecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(handler)
	
		handler.ServeHTTP(resrecorder, request)
			
		assert.Equal(t, resrecorder.Code, http.StatusOK, "Unable to render webpage returned status code: %v expected: %v", resrecorder.Code, http.StatusOK)
}
