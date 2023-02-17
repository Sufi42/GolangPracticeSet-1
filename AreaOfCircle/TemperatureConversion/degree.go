package main

import "fmt"

func main() {
	fmt.Println("Enter temperature in Degree celcius")
	var temp float32
	fmt.Scanln(&temp)
	fareheit := ((temp * 9) / 5) + 32
	fmt.Println(fareheit)
}
