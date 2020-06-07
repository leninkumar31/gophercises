package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"../../cyoa"
	"../../handler"
)

func main() {
	jsonfile := flag.String("json", "gophers.json", "Json file which has stories")
	htmlfile := flag.String("html", "layout.html", "HTML to render stories")
	flag.Parse()
	// open json file which contains stories
	file := openFile(*jsonfile)
	story, err := cyoa.JsonStory(file)
	if err != nil {
		exit("Not able to read file")
	}
	// create template using htmlfile
	t := template.Must(template.ParseFiles(*htmlfile))
	handler := handler.DefaultHandler(t, story)
	fmt.Println("Server started listening on port 8080...")
	http.ListenAndServe(":8080", handler)
}

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		exit("Unable to open file")
	}
	return file
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
