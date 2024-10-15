package main

import (
	"fmt"
	"net/http"
	"scraper/api"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html") // Serve the HTML file
}

func main() {
	// Serve the index.html at /home
	http.HandleFunc("/home", serveHome)

	// Serve the API endpoint for scraping data
	http.HandleFunc("/scrape-data", api.ApiHandler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
