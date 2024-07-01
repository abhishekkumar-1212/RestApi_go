package main


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// here create json by using the struct
type Todo struct {
	Pizza_Price    int    `json:"pizza price"`
	Flavour        string   `json:"flavour"`
	Brand    string `json:"brand"`
	Payment_Status bool   `json:"Payment Status"`
}

func performPostRequest() {
	todo := Todo{
		Pizza_Price:    593,
		Flavour: "Chesse Burst",
		Brand:     "Dominos",
		Payment_Status: true,
	}
	// when we send data in the form struct to JSon
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error Marshalling: ", err)
		return
	}

	// But we pass data as the string :
	// Convert JsonData to string
	jsonString := string(jsonData)

	// 1:Now data is ready :
	// 2: Now we Call
	// convert string data to reader :
	jsonReader := strings.NewReader(jsonString)

	myURl := "https://jsonplaceholder.typicode.com/todos"

	// Send Post request
	res, err := http.Post(myURl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}
	defer res.Body.Close()

	// Now we convert response in the readable form :
	// Now read the response
	data, _ := ioutil.ReadAll(res.Body)
fmt.Println("response: ", string(data));


}

// It is the main function : 
func main() {
	fmt.Println("Welcome to get request... ")
// Now Call the function :
performPostRequest()
}
