package main

import "fmt"

func main() {
	fmt.Println("If else in golang: ")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular user: "
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 length count"

	}

	fmt.Println(result)

// check odd and even in the golang 
if 10%2==0{
	fmt.Println("EVEN")
} else{
	fmt.Println("ODD")
}

// This type we  generally used in during web_request and assign in the go
if num:=3; num<10{
  fmt.Println("Num is Less than 10: ")
} else{
	fmt.Println("Num is greater than 10")

}

// If not equals to null
/*
 if err!=nill{

 }

*/

}


