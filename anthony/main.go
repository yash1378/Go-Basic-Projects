// Go program to illustrate how to
// find the capacity of the channel

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Email   string   `json:"email"`
	Phone   string   `json:"phone"`
	Hobbies []string `json:"hobbies"`
}

func main() {
	p := Person{
		Name:  "John Jones",
		Age:   26,
		Email: "johnjones@email.com",
		Phone: "89910119",
		Hobbies: []string{
			"Swimming",
			"Badminton",
		},
	}

	b, err := json.MarshalIndent(p, "", "  ") // Marshal the struct to JSON
	if err != nil {
		log.Fatalf("Unable to marshal due to %s\n", err)
	}

	fmt.Println(string(b)) // Print the JSON output
}
