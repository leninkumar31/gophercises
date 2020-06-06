package handler

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

type urldata []struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func YAMLHandler(yamldata []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// parse yaml data to struct
	var paths urldata
	err := yaml.Unmarshal(yamldata, &paths)
	if err != nil {
		return nil, err
	}
	// convert array of structs to map of strings to strings
	pathToURLMap := make(map[string]string)
	for _, path := range paths {
		pathToURLMap[path.Path] = path.URL
	}
	return MapHandler(pathToURLMap, fallback), nil
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if url, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, url, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	}
}
