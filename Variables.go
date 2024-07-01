package main

import "fmt"  // fmt is for variables , datatypes values 

// const LoginToken string="adfs" // Login 'L' indicate the public in nature 



func main() {
	fmt.Println("Variables")
	var username string = "Hitesh"

	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

var isLoggedIn bool=false;
fmt.Println(isLoggedIn)
fmt.Printf("Variable is of type: %T \n",isLoggedIn)


var smallVal uint8 =255;
fmt.Println(smallVal)
fmt.Printf("Variable is of type: %T \n",smallVal)


var smallFloat float32 =255.1234;
fmt.Println(smallFloat)
fmt.Printf("Variable is of type: %T \n",smallFloat)

// default values and some aliases: ->
var anotherVariable int
fmt.Println(anotherVariable)
fmt.Printf("Variable is of type: %T \n" ,anotherVariable)

// implicit type to declare the variable:

var website="www.google.com"
fmt.Println(website)

// no variable style 
numberOfUser:=2345324.0  // := is called the volurus operator
fmt.Println(numberOfUser)  // 





}
