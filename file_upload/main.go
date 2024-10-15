package main

import (
	"log"
	"net/http"
	"upload/controllers"
)

func main() {
	// Serve static files from the 'controllers/static' directory
	fs := http.FileServer(http.Dir("./controllers/static")) // Use relative path to 'controllers/static'
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fs)) // Strip the '/static/' prefix for serving static files

	// Serve the index file
	mux.HandleFunc("/", controllers.Index)
	mux.HandleFunc("/upload", controllers.Upload)

	// Start the server
	if err := http.ListenAndServe(":4500", mux); err != nil {
		log.Fatal(err)
	}
}
