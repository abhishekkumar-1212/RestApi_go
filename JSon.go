package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}


func main() {

	fmt.Println("Learning json")
	person := Person{Name: "John", Age: 34, IsAdult: true}
	fmt.Println("Person data is : ", person)

	// Convert person into JSON ENcoding (Marshalling )
	jsonData, err := json.Marshal(person)

if err!=nil{
fmt.Println("Error marshalling ", err)
return
}

fmt.Println("Person Data is: ", string(jsonData))

// Decoding( Unmarshalling )
var personData Person

err=json.Unmarshal(jsonData, &personData)

if err!=nil{
fmt.Println("Error unmarshalling ", err)

}
fmt.Println("Person data is: ", personData)

}
