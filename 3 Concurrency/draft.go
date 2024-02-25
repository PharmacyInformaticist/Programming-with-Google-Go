package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Create an unsorted array of integers
	//arr := []int{5, 3, 8, 1, 7, 2, 6, 4}
	arr := takeInput()
	// Create channels for communication between goroutines
	c1 := make(chan []int)
	c2 := make(chan []int)
	c3 := make(chan []int)
	c4 := make(chan []int)
	c5 := make(chan []int)

	// Split the array into four sub-arrays and sort each sub-array using a separate goroutine
	go sortSubArray(arr[:len(arr)/4], c1)
	go sortSubArray(arr[len(arr)/4:len(arr)/2], c2)
	go sortSubArray(arr[len(arr)/2:3*len(arr)/4], c3)
	go sortSubArray(arr[3*len(arr)/4:], c4)

	// Merge the sorted sub-arrays into a single sorted array using four goroutines
	go merge(c1, c2, c5)
	go merge(c3, c4, c5)
	go merge(c5, c5, c5)

	// Wait for the final sorted array to be produced and print it
	fmt.Println(<-c5)
}

func sortSubArray(arr []int, c chan []int) {
	// Sort the sub-array
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	// Send the sorted sub-array to the channel
	c <- arr
}

func merge(c1, c2, c3 chan []int) {
	// Wait for the two sorted sub-arrays to be received from the channels
	arr1 := <-c1
	arr2 := <-c2

	// Merge the two sorted sub-arrays into a single sorted array
	merged := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged[k] = arr1[i]
			i++
		} else {
			merged[k] = arr2[j]
			j++
		}
		k++
	}
	for i < len(arr1) {
		merged[k] = arr1[i]
		i++
		k++
	}
	for j < len(arr2) {
		merged[k] = arr2[j]
		j++
		k++
	}

	// Send the merged sorted sub-array to the channel
	c3 <- merged
}

// func takes input and creates a []int
func takeInput() []int {
	var int_slc []int
take_input:
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("enter a space-separated sequence of integers.")
	scanner.Scan()
	inp := strings.ToLower(scanner.Text())
	req := strings.Split(inp, " ")
	for _, v := range req {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Printf("[Error] Invalid input!")
			goto take_input
		}
		int_slc = append(int_slc, i)
	}
	fmt.Println("The unsorted slice of integers is: ", int_slc)
	return int_slc
}
