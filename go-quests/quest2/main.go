package main

import (
	"fmt"
	"net/http"
	"flag"
	"io/ioutil"
	"./urlshort"
)

func main() {
	mux := defaultMux()

	// Set up flags for command line
	yamlFlag := flag.String("f","", "a yaml file to read url mappings in")
	flag.Parse()

	data, err := ioutil.ReadFile(*yamlFlag)
	if err != nil {
		fmt.Println("Yaml file reading error", err)
		return
	}

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
// 	yaml := `
// - path: /urlshort
//   url: https://github.com/gophercises/urlshort
// - path: /urlshort-final
//   url: https://github.com/gophercises/urlshort/tree/solution
//`
	yamlHandler, err := urlshort.YAMLHandler([]byte(data), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}