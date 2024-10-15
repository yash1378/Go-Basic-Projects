package controllers

// Post struct represents the data you get from scraping the JSON API.
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// Comment struct represents the data you get from scraping the JSON API.
type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

// Result struct represents the data you get from scraping a website.
type Result struct {
	Posts    []Post
	Comments []Comment
	Error    error
}
