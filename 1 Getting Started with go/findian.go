package main

import (
	"fmt"
	"log"
	"strings"
)

// string var to be evaluated
var (
	str    string
	result bool
	_      string
)

// contains function evaluates bool
func str_check(str string) {
init_check:
	fmt.Println("Enter a string to evaluate: ")
	_, err := fmt.Scan(&str)
	if err != nil {
		log.Printf("[Error] Invalid input!")
		goto init_check
	}
	result = ((strings.HasPrefix(str, "i") || strings.HasPrefix(str, "I")) && (strings.Contains(str, "a") || strings.Contains(str, "A")) && (strings.HasSuffix(str, "n") || strings.HasSuffix(str, "N")))
}

// main func loops ask for str input until user finishes execution
func main() {
	var ans string = "c"
	for {
		str_check(str)
		if result == true {
			fmt.Println("Found!")
			fmt.Println("Press c to continue or press any other key to finish")
			fmt.Scan(&ans)
		} else {
			fmt.Println("Not Found!")
			fmt.Println("Press c to continue or press any other key to finish")
			fmt.Scan(&ans)
		}
		if ans != "c" {
			fmt.Println("Program terminated")
			break
		}
	}
}
