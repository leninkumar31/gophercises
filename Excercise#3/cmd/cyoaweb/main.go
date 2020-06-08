package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"../../cyoa"
	"../../handler"
	"../../utils"
)

func main() {
	port := flag.Int("port", 8080, "Port on choose your own adventure runs")
	jsonfile := flag.String("json", "gophers.json", "Json file which has stories")
	htmlfile := flag.String("html", "temp.html", "HTML to render stories")
	flag.Parse()
	// open json file which contains stories
	file := utils.OpenFile(*jsonfile)
	story, err := cyoa.JsonStory(file)
	if err != nil {
		utils.Exit("Not able to read file")
	}
	// create template using htmlfile
	t := template.Must(template.ParseFiles(*htmlfile))
	// handler := handler.DefaultHandler(t, story)
	mux := http.NewServeMux()
	h := handler.NewHandler(story, handler.WithTemplate(t), handler.WithPathParseFunc(pathParseFunc()))
	mux.Handle("/story/", h)
	fmt.Printf("Server started listening on port :%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func pathParseFunc() func(r *http.Request) string {
	return func(r *http.Request) string {
		path := r.URL.Path
		if path == "/story" || path == "/story/" {
			path = "/story/intro"
		}
		return path[len("/story/"):]
	}
}
