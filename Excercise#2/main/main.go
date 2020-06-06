package main

import (
	"fmt"
	"net/http"
	"os"

	"../handler"
)

type urldata []struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func main() {
	// create a default multiplexer
	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handler.MapHandler(pathsToUrls, mux)
	yamldata := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlhandler, err := handler.YAMLHandler([]byte(yamldata), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server listening at 8080...")
	http.ListenAndServe(":8080", yamlhandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
