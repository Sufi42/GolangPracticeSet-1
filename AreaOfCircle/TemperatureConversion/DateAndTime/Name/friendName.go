package main

import "fmt"

func main() {
	fmt.Println("Enter friend first name : ")
	var firstName string
	fmt.Scanln(&firstName)
	fmt.Println("Enter Last NAme of Friend : ")
	var lastName string
	fmt.Scanln(&lastName)
	fmt.Println(firstName, lastName)
}
