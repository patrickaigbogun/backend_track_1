package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"regexp"
	"strconv"
	// "strings"
)

func checkForArmstrong(number int) bool {
	// Convert the number to a string to easily extract digits
	numStr := fmt.Sprintf("%d", number)
	numDigits := len(numStr)
	sum := 0
	tempNumber := number

	// Loop through each digit and raise it to the power of numDigits
	for tempNumber > 0 {
		digit := tempNumber % 10
		sum += int(math.Pow(float64(digit), float64(numDigits)))
		tempNumber /= 10
	}

	// Check if the sum is equal to the original number
	return sum == number
}

func checkForProperties(number int) interface{} {
	var result interface{}

	a := [2]string{"armstrong", "odd"}
	b := [2]string{"armstrong", "even"}
	c := [1]string{"odd"}
	d := [1]string{"even"}

	x := number % 2

	switch {
	case x != 0 && checkForArmstrong(number):
		result = b
	case x == 0 && checkForArmstrong(number):
		result = a
	case x != 0:
		result = c
	default:
		result = d
	}

	return result
}

func checkForSum(number int) int {
	sum := 0
	for number > 0 {
		sum += number % 10
		number /= 10
	}
	return sum
}

// fetchFunFact retrieves a fun fact about the number from the numbersapi.com API
func fetchFunFact(number int) string {
	// Prepare the URL for the API request
	url := fmt.Sprintf("http://numbersapi.com/%d/math", number)

	// Make the GET request to numbersapi.com
	resp, err := http.Get(url)
	if err != nil {
		// If an error occurs while making the request, return an error message
		return "Error fetching fun fact"
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Error reading response body"
	}

	// Convert the response body to a string
	funFact := string(body)

	return funFact

	// If there's no fun fact in the response, return a default message
	// return "No fun fact found"
}

func checkForPerfectNumber(number int) string {
	sum := 0
	rem := 0

	for count := 1; count < number; count++ {
		rem = number % count
		if rem == 0 {
			sum = sum + count
		}
	}

	if number == sum {
		return "true"
	} else {
		return "false"
	}
}

func checkForPrime(number int) string {
	// Edge case: numbers less than or equal to 1 are not prime
	if number <= 1 {
		return "false"
	}

	// Check for divisibility from 2 to the square root of the number
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			return "false"
		}
	}

	// If no divisor is found, it's a prime number
	return "true"
}

func handleNumber(w http.ResponseWriter, r *http.Request) {
	// Define bad request response
	badRequestResponse := map[string]string{"error": "true", "message": "alphabet"}

	// Marshaling bad request response for later use
	badRequestResponseInJSON, err := json.Marshal(badRequestResponse)
	if err != nil {
		http.Error(w, "Error marshaling bad request response", http.StatusInternalServerError)
		return
	}

	// Only allow GET requests (method check)
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET methods are allowed", http.StatusMethodNotAllowed)
		return
	}

	q := r.URL.Query()
	numberParam := q.Get("number")

	// Check if the 'number' parameter is alphabetic
	isAlphabet := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(numberParam)
	if isAlphabet {
		// If it's an alphabet, return the bad request response
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(badRequestResponseInJSON)
		return
	}

	// Convert the number parameter to an integer
	// If it's a valid number, this will succeed
	number, err := strconv.Atoi(numberParam)
	if err != nil {
		http.Error(w, "Invalid number parameter", http.StatusBadRequest)
		return
	}

	// Respond with the received number
	fmt.Printf("Received number: %d\n", number)

	// Calculate properties
	isPrime := checkForPrime(number)
	isPerfect := checkForPerfectNumber(number)
	properties := checkForProperties(number)
	digitSum := checkForSum(number)
	funFact := fetchFunFact(number)

	// Prepare the response
	response := map[string]interface{}{
		"number":     number,
		"is_prime":   isPrime,
		"is_perfect": isPerfect,
		"properties": properties,
		"digit_sum":  digitSum,
		"fun_fact":   funFact,
	}

	// Marshal the response to JSON
	responseInJson, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}

	// Set headers and send response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(responseInJson)
}

func main() {
	// Log server startup
	log.Println("Server started on port 7000")

	// Handle the "/api/clasify-number" endpoint
	http.HandleFunc("/api/classify-number", handleNumber)

	// Start the server
	log.Fatal(http.ListenAndServe(":7000", nil))
}
