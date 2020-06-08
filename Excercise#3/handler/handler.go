package handler

import (
	"html/template"
	"log"
	"net/http"

	"../cyoa"
)

var tmp *template.Template
var defaultPathParseFunc func(r *http.Request) string

func init() {
	tmp = template.Must(template.ParseFiles("layout.html"))
	defaultPathParseFunc = DefaultPathParseFunc
}

type HandlerOpts func(h *handler)

func WithTemplate(t *template.Template) HandlerOpts {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathParseFunc(fn func(r *http.Request) string) HandlerOpts {
	return func(h *handler) {
		h.pathFunc = fn
	}
}

func NewHandler(s cyoa.Story, opts ...HandlerOpts) http.Handler {
	h := handler{story: s, t: tmp, pathFunc: defaultPathParseFunc}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type handler struct {
	story    cyoa.Story
	t        *template.Template
	pathFunc func(r *http.Request) string
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFunc(r)
	if chapter, ok := h.story[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found...", http.StatusNotFound)
}

func DefaultPathParseFunc(r *http.Request) string {
	path := r.URL.Path
	if path == "" || path == "/" {
		path = "/intro"
	}
	return path[1:]
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
