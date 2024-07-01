package main

import (
	"bufio" // It includes for input and output 
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input: "

	fmt.Println(welcome)
	fmt.Printf("The Type of this: %T \n", welcome)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating of the Pizza: ")

	// Comma ok / err ok (It is used for error handling and try and catch )
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating ", input)

}
