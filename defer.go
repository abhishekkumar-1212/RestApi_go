package main
import "fmt"


// defers works on last in first out [ LIFO ]

func main(){
fmt.Println("Hello")
defer fmt.Println("One")
defer fmt.Println("Two")
defer fmt.Println("Three")
myDefer()
}

func myDefer(){

// reverse print the loop by using defer 
for i:=0 ;i<5;i++{
	defer fmt.Println(i)
}

}