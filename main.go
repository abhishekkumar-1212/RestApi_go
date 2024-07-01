package main
import "fmt"

func main(){

fmt.Println("WELCOME TO MAIN FUNCTION")
greet()

ans:= adder(3,5)
fmt.Println("Answer is: ", ans)

// try to use the pro function:
var proRes int
proRes=proAdder(2,5,8,7,3)
fmt.Println("Pro result is:",proRes)

}

func greet(){
fmt.Println("Namastay")

}

// add two values by using the function
func adder(valone int , valtwo int) int  {  // what kind of value we returned
return valone+valtwo
}

// If user pass any value: inside the function
func proAdder(values... int ) int {
total:=0

// iterating in golang: 
for _,val:=range values{
	total+=val

}
return total
}