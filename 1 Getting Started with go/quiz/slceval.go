package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input string
	var slice = make([]int, 0, 3)

	// run a inifinite for loop until "X" is provided as input
	for {
		// request for an integer via the prompt
		fmt.Println("Please enter a integer value:")
		fmt.Scan(&input)
		// check if it is a stop criteria
		if strings.ContainsAny(strings.ToLower(input), "x") {
			break
		} else {
			// convert string to int
			var val, err = strconv.Atoi(input)
			// check if there is any error
			if err != nil {
				break
			}
			// and append the new value to the slice
			slice = append(slice, val)
			// sort the slice using the sort module
			sort.Ints(slice)
			// print the slice to console
			fmt.Println("The current slice is: ", slice, "")
		}
	}
}
