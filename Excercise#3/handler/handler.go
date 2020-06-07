package handler

import (
	"html/template"
	"net/http"

	"../cyoa"
)

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
