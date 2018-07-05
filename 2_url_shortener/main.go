package main

import (
	"io/ioutil"
	"fmt"
	"gopkg.in/yaml.v2"
	"net/http"
	"flag"

)

var ymlPath string

func main() {
	flag.StringVar(&ymlPath, "yml", "urls.yml", "the path to a .yml file that contains paths and urls")
	flag.Parse()

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/hi":   "https://godoc.org/github.com/gophercise/urlshort",
		"/yaml": "https://godoc.org/gopkg.in/yaml.v2",
	}

	mhandler := MapHandler(pathsToUrls, mux)

	yml, err := ioutil.ReadFile(ymlPath)

	yhandler, err := YAMLHandler([]byte(yml), mhandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yhandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(yml)

	if err != nil {
		return nil, err
	}
	yaml.Unmarshal(yml, &pathUrls)
	pathsToUrls := buildMap(pathUrls)
	return MapHandler(pathsToUrls, fallback), nil
}

func buildMap(pathUrls []pathUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.Url
	}
	return pathsToUrls
}

func parseYaml(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}
