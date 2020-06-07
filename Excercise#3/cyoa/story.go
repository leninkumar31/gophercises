package cyoa

import (
	"encoding/json"
	"io"
)

func JsonStory(r io.Reader) (Story, error) {
	story := make(Story)
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&story)
	if err != nil {
		return nil, err
	}
	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text        string `json:"text"`
	ChapterName string `json:"arc"`
}
