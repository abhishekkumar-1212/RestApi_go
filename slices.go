package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices that is most uses in go ")

	var fruitList = []string{"Apple ", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = fruitList[1:3] //here we used the slice method

	fmt.Println(fruitList)

	// making an array alloted to the memory
	highscores := make([]int, 4)
	highscores[0] = 4
	highscores[1] = 24
	highscores[2] = 34
	highscores[3] = 2
	fmt.Println(highscores)

	// now add the highscores and append
	highscores = append(highscores, 324, 45, 64)
	fmt.Println(highscores)

	// here we sorts the integers values:
	sort.Ints(highscores)

	// to check if the given array is sorted or not.
	fmt.Println("Is the array is sorted:" , sort.IntsAreSorted(highscores))

// how to remove value from slices based on index

var courses=[] string{"ReactJs", "Javascript", "Swift ","Python"}
fmt.Println(courses)

var index int =2;
courses=append(courses[:index],courses[index+1:]...)
fmt.Println(courses)


  



}
