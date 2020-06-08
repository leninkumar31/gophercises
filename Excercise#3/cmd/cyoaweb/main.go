package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"../../cyoa"
	"../../handler"
	"../../utils"
)

func main() {
	port := flag.Int("port", 8080, "Port on choose your own adventure runs")
	jsonfile := flag.String("json", "gophers.json", "Json file which has stories")
	// htmlfile := flag.String("html", "layout.html", "HTML to render stories")
	flag.Parse()
	// open json file which contains stories
	file := utils.OpenFile(*jsonfile)
	story, err := cyoa.JsonStory(file)
	if err != nil {
		utils.Exit("Not able to read file")
	}
	// create template using htmlfile
	// t := template.Must(template.ParseFiles(*htmlfile))
	// handler := handler.DefaultHandler(t, story)
	h := handler.NewHandler(story)
	fmt.Printf("Server started listening on port :%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
