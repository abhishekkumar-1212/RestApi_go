package main
import "fmt"

func main(){
fmt.Println("Welcome to array ")

// in go array is less use 
var fruitList[4] string ;
fruitList[0]="Apple"
fruitList[1]="Tomato"
fruitList[3]="Peach"
fmt.Println("The fruitList is : ", fruitList)
fmt.Println("Length of List is: ", len(fruitList))

// second is  vegList 
var vegList = [3]string{"Bangan ","Adrak","Chilli"}
fmt.Println("Veg List is: ",vegList)
fmt.Println("Length of vegList is : ",len(vegList))

}