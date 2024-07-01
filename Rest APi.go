package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"Completed "`
}

func performGetRequest() {

	res, err := http.Get("http://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error getting: ", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 { // or we used http.StatusOk
		fmt.Println("Error in getting response: ", res.Status)
	}

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading: ", err)
	// 	return
	// }

	// fmt.Println("DATA: ",string(data) )

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding: ", err)
		return
	}
	fmt.Println("Todo: ", todo)
}

// here we send the Data with the help of Post method
func performPostRequest() {
	todo := Todo{
		UserID:    56,
		Title:     "INtern bros",
		Completed: true,
	}

	// Convert the Todo struct to JSON :
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling: ", err)
		return
	}

	// COnvert json data to string :
	jsonString := string(jsonData)

	// Convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos/"

	// send the POST Request
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}

	defer res.Body.Close()

	// Convert response body to readable form
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response: ", err)
		return
	}
	fmt.Println("Response: ", string(data))
fmt.Println("Response status: ", res.Status )

}

func main() {
	fmt.Println("Learning CRUD...")
	// Now call the functions

	//performGetRequest()
	performPostRequest()

}
