package main

import (
	"fmt"
	"log"
)

// float variable to be truncated
var x float64

// truncate function returns int number
func truncated(x float64) int64 {
outside:
	fmt.Println("Enter a number to truncate: ")
	_, err := fmt.Scan(&x)
	if err != nil {
		log.Printf("[Error] Invalid input!")
		goto outside
	}
	x1 := int64(x)
	return x1
}

// main func loops ask for number input until user finishes execution
func main() {
	var ans string = "c"
	for {
		fmt.Println("The truncated number is", truncated(x))
		fmt.Println("Press c to continue or press any other key to finish")
		fmt.Scan(&ans)
		if ans != "c" {
			fmt.Println("Program terminated")
			break
		}
	}
}
