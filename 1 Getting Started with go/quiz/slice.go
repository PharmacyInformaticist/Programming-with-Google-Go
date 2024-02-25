package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func main() {
	var int_slc []int
	for {
	take_input:
		fmt.Println("This programs creates a sorted slice of integers.")
		fmt.Println("Enter an integer(s) or Press x to exit the program : ")
		var inp string
		fmt.Scan(&inp)
		if inp == "x" || inp == "X" {
			fmt.Println("Program terminated")
			fmt.Println("The final sorted slice is: ", int_slc)
			break
		} else {
			i, err := strconv.Atoi(inp)
			if err != nil {
				log.Printf("[Error] Invalid input!")
				goto take_input
			}
			int_slc = append(int_slc, i)
			sort.Ints(int_slc)
			fmt.Println("The sorted slice of integers is: ", int_slc)
		}
	}
}
