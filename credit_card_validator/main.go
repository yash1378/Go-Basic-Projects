package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// RequestBody represents the expected JSON payload structure.
type RequestBody struct {
	CreditCardNumber string `json:"creditCardNumber"`
}

// ResponseBody represents the structure of the JSON response.
type ResponseBody struct {
	IsValid bool   `json:"isValid"`
	Message string `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}

// Luhn algorithm implementation to validate a credit card number.
func luhnAlgorithm(cardNumber string) bool {
	sum := 0
	double := false

	// Traverse the card number from right to left
	for i := len(cardNumber) - 1; i >= 0; i-- {
		n, err := strconv.Atoi(string(cardNumber[i]))
		if err != nil {
			return false // Invalid character in card number
		}

		if double {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		double = !double
	}

	return sum%10 == 0
}

func indexhandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to plain text
	w.Header().Set("Content-Type", "text/plain")

	// Set a custom status code (optional, defaults to 200 OK if not set)
	w.WriteHeader(http.StatusOK)

	response := Response{Message: "this is the reponsed"}
	// Marshal the response struct into JSON
	// jsonResponse, err := json.Marshal(response)
	// if err != nil {
	// 	http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	// 	return
	// }

	json.NewEncoder(w).Encode(response)
	// Write the JSON response body
	// w.Write(jsonResponse)

	// Write the response body
	// w.Write([]byte("Hello, World!"))
}

func validate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, " Only Post method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var rbody RequestBody

	if err := json.NewDecoder(r.Body).Decode(&rbody); err != nil {
		http.Error(w, "Invalid json payload", http.StatusBadRequest)
		return
	}

	if rbody.CreditCardNumber == "" {
		http.Error(w, "Credit Card Number is Missing", http.StatusBadRequest)
		return
	}

	isValid := luhnAlgorithm(rbody.CreditCardNumber)

	response := ResponseBody{
		IsValid: isValid,
		Message: "Credit Card Validation complete",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexhandler)
	mux.HandleFunc("/validate", validate)
	// Start the HTTP server on port 8080
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", mux)
}
