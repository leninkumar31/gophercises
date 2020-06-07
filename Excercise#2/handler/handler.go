package handler

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

type urldata []struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// YAMLHandler ......
func YAMLHandler(yamldata []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// parse yaml data to struct
	pathsToUrls, err := parseYAML(yamldata)
	if err != nil {
		return nil, err
	}
	// convert array of structs to map of strings to strings
	pathToURLMap := buildMap(pathsToUrls)
	return MapHandler(pathToURLMap, fallback), nil
}

// MapHandler .....
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if url, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, url, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func buildMap(pathsToUrls urldata) map[string]string {
	pathToURLMap := make(map[string]string)
	for _, path := range pathsToUrls {
		pathToURLMap[path.Path] = path.URL
	}
	return pathToURLMap
}

func parseYAML(data []byte) (urldata, error) {
	var paths urldata
	err := yaml.Unmarshal(data, &paths)
	if err != nil {
		return nil, err
	}
	return paths, nil
}
