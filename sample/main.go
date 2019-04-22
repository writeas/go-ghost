package main

import (
	"fmt"
	"github.com/writeas/go-ghost"
	"github.com/writeas/go-ghost/admin"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: sample GHOST_URL API_KEY\n")
		os.Exit(1)
	}

	url := os.Args[1]
	apiKey := os.Args[2]

	c := ghost.NewClient(url, apiKey)
	err := admin.AddPost(c, ghost.PostParams{
		Title:    ghost.String("Posting via Go"),
		Markdown: ghost.String(`This is a **test post** made with the [go-ghost](https://github.com/writeas/go-ghost) library.`),
		Status:   ghost.String("published"),
	})
	if err != nil {
		fmt.Printf("AddPost: %v\n", err)
	}
}
