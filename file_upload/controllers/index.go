package controllers

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	// Correct the path to the 'index.html' relative to where the program is run
	http.ServeFile(w, r, "controllers/static/index.html") // Correct the path here
}
