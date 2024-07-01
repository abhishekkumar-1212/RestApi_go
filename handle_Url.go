package main

import (
	"fmt"                    
	"net/url"               
)

/*
Scheme:  https
Host:  www.google.com (Domain Name)
Path:
RawQuery:
*/

// SHORTCUT :->
func main() {
	fmt.Println("Handle Url")
	myUrl := "https://www.google.com"
	fmt.Printf("Type of URL: %T\n", myUrl)

	// string converted into url(url.Parse)
	parsedUrl, err := url.Parse(myUrl)
if err!=nil{
	fmt.Println("Can't Parse URL",err)
}
// One we convert string into URL the given myUrl then we have different components in go . 
fmt.Printf("Type of URL: %T\n",parsedUrl)
// 1: Scheme
fmt.Println("Scheme: ",parsedUrl.Scheme)
//2: Host 
fmt.Println("Host: ",parsedUrl.Host)
// 3: Path 
fmt.Println("Path: ", parsedUrl.Path)
//4: RawQuery : 
fmt.Println("RawQuery: ",parsedUrl.RawQuery)
 
// we also modify the Url Components
// In industrial level we have existing Url (That existing Url is changes to make the new Url  )

parsedUrl.Host="www.facebook.com"
fmt.Println("New Host is: ",parsedUrl.Host)

// Constructing Url string from Url object 
newUrl:=parsedUrl.String() // string 'S' is capital in go 
fmt.Println("new URL: ", newUrl)

}
