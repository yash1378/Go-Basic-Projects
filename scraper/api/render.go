package api

import (
	"encoding/json"
	"net/http"
	"scraper/controllers"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	postCh := make(chan []controllers.Post)
	commentCh := make(chan []controllers.Comment)

	go controllers.ScrapePosts("https://jsonplaceholder.typicode.com/posts", postCh)
	go controllers.ScrapeComments("https://jsonplaceholder.typicode.com/comments", commentCh)

	posts := <-postCh
	comments := <-commentCh

	if posts == nil || comments == nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	// Create the final result struct with both posts and comments
	result := controllers.Result{
		Posts:    posts,
		Comments: comments,
	}

	// Send the result as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
