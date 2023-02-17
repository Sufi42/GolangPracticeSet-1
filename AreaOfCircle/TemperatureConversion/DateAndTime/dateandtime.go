package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	var (
		day   int = currentTime.Day()
		month int = int(currentTime.Month())
		year  int = currentTime.Year()
	)
	fmt.Println("Current time is : ", currentTime)
	fmt.Println("Day is : ", day)
	fmt.Println("Month of year is  : ", month)
	fmt.Println("Year is : ", year)
	fmt.Println("DD-MMM_YYYY : ", day, "-", month, "-", year)

}
