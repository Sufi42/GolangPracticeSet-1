package main

import "fmt"

func main() {
	fmt.Println("Enter diameter of circle")
	var diameter float32
	const pie = 3.14
	fmt.Scanln(&diameter)
	radius := diameter / 2
	area := pie * radius * 2
	fmt.Println(area)

}
