package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web services: ")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
 if err!=nil{
	panic(err)
 }
 defer res.Body.Close()  // Active resources which could not in use we free that by using "defer " keyword
 fmt.Printf("Type of response:%T\n",res)

// Read the response body
data,err:=ioutil.ReadAll(res.Body)
if err!=nil{
	panic(err)
}
// data in the form of array of byte .
fmt.Println("response: ",string(data))



}
