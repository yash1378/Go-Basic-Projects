package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mgutz/ansi"
)

func printPost(post Post) {
	// Print the post in green
	fmt.Printf(ansi.Color(fmt.Sprintf("Post ID: %d - Title: %s\n", post.ID, post.Title), "green"))
}

func ScrapePosts(url string, ch chan<- []Post) {
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

	var posts []Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		ch <- nil
		return
	}

	ch <- posts
}
