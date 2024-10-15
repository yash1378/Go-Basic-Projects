package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mgutz/ansi"
)

func printComment(comment Comment) {
	// Print the comment in red
	fmt.Printf(ansi.Color(fmt.Sprintf("Comment ID: %d - Name: %s\n", comment.ID, comment.Name), "red"))
}

func ScrapeComments(url string, ch chan<- []Comment) {
	time.Sleep(2 * time.Second)

	resp, err := http.Get(url)
	if err != nil {
		ch <- nil
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ch <- nil
		return
	}

	var comments []Comment
	if err := json.NewDecoder(resp.Body).Decode(&comments); err != nil {
		ch <- nil
		return
	}

	ch <- comments
}
