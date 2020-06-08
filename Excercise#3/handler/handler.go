package handler

import (
	"html/template"
	"net/http"

	"../cyoa"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseFiles("layout.html"))
}

func NewHandler(s cyoa.Story) http.Handler {
	return handler{story: s}
}

type handler struct {
	story cyoa.Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tmp.Execute(w, h.story["intro"])
	if err != nil {
		panic(err)
	}
}

func DefaultHandler(t *template.Template, story cyoa.Story) http.Handler {
	mux := http.NewServeMux()
	for key, val := range story {
		mux.HandleFunc("/"+key, getHandler(t, val))
	}
	return mux
}

func getHandler(t *template.Template, chapter cyoa.Chapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, chapter)
	}
}
