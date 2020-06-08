package utils

import (
	"fmt"
	"os"
)

func OpenFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		Exit("Unable to open file")
	}
	return file
}

func Exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
