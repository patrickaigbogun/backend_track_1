package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handleNumber(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET methods are allowed", http.StatusMethodNotAllowed)
		return
	}

	q := r.URL.Query()
	numberParam := q.Get("number")

	// Convert the number parameter to an integer
	number, err := strconv.Atoi(numberParam)
	if err != nil {
		http.Error(w, "Invalid number parameter", http.StatusBadRequest)
		return
	}

	// Respond with the received number
	fmt.Printf("Received number: %d\n", number)

	// Send a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	response := map[string]interface{}{"received_number": number}
	responseInJson, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}

	w.Write(responseInJson)
}

func main() {
	// Log server startup
	log.Println("Server started on port 7000")

	// Handle the "/api/clasify-number" endpoint
	http.HandleFunc("/api/clasify-number", handleNumber)

	// Start the server
	log.Fatal(http.ListenAndServe(":7000", nil))

}
