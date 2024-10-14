package main

import (
	"encoding/json"
	"fmt"
	"io"
	"learningpackage/simple"

	// "log"
	"net/http"
)

type post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func hello(w http.ResponseWriter, req *http.Request) {
	// Write a plain text response
	w.Write([]byte("Hello, world!"))
	// fmt.Fprintf(w, "Hello, %s!", "world")

	// // Using io.WriteString
	// io.WriteString(w, " Hello again!")

	apiurl := "https://jsonplaceholder.typicode.com/posts"

	resp, err := http.Get(apiurl)
	defer resp.Body.Close()
	if err != nil {
		http.Error(w, "Failed to fetch data from the API", http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read API response", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("API request failed with status: %s", resp.Status), http.StatusInternalServerError)
		return
	}

	var apiResponse []post
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		http.Error(w, "Failed to parse API response", http.StatusInternalServerError)
		return
	}
	var responseString string
	for _, post := range apiResponse {
		responseString += fmt.Sprintf(
			"UserID: %d\nID: %d\nTitle: %s\nBody: %s\n\n",
			post.UserID,
			post.ID,
			post.Title,
			post.Body,
		)
	}

	// Write the formatted string to the HTTP response
	fmt.Fprintf(w, responseString)

}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func multiple(a int, b string, c bool) (bool, float32) {
	fmt.Println(a, b, c)
	return true, 3.45
}

var p, r, t = 5000.0, 10.0, 1.0

/*
* init function to check if p, r and t are greater than zero
 */
func init() {
	fmt.Println("Main package initialized")
	// if p < 0 {
	// 	log.Fatal("Principal is less than zero")
	// }
	// if r < 0 {
	// 	log.Fatal("Rate of interest is less than zero")
	// }
	// if t < 0 {
	// 	log.Fatal("Duration is less than zero")
	// }
}

type currency struct {
	name   string
	symbol string
}

func main() {
	simple.Calculate()
	fmt.Println("from the main function")
	// a := new(int)
	// fmt.Println(a, *a)
	// *a = 101
	// fmt.Println(a, *a)
}
