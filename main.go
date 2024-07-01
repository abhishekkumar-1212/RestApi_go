package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Define the Shop struct
type Shop struct {
	Pizza_Price    int    `json:"pizza price"`
	Flavour        string `json:"flavour"`
	Brand          string `json:"brand"`
	Payment_Status bool   `json:"Payment Status"`
}

// Function to perform POST request
func performPostRequest() {
	todo := Shop{
		Pizza_Price:    593,
		Flavour:        "Cheese Burst",
		Brand:          "Dominos",
		Payment_Status: true,
	}
	// Convert struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error Marshalling: ", err)
		return
	}

	// Convert JSON data to string
	jsonString := string(jsonData)

	// Convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos"

	// Send POST request
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}
	defer res.Body.Close()

	// Read the response
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response: ", string(data))
}

// Function to perform PUT request
func performUpdateRequest() {
	x := Shop{ // Updated data
		Pizza_Price:    600,
		Flavour:        "Veg Paneer",
		Brand:          "McDonald's",
		Payment_Status: true,
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(x)
	if err != nil {
		fmt.Println("Error in Marshalling", err)
		return
	}

	// Convert JSON data to string
	jsonString := string(jsonData)

	// Convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	const myURL = "https://jsonplaceholder.typicode.com/todos/1"

	// Create PUT request
	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT request: ", err)
		return
	}

	// Set the content type
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}
	defer res.Body.Close()

	// Read the response
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response: ", string(data))
}

// Main function
func main() {
	fmt.Println("Welcome to Update request... ")
	// Uncomment the following line to perform a POST request
	// performPostRequest()
	performUpdateRequest()
}
