package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

// func that sorts slice an sends to channel

func sortArray(s []int, c chan []int) {
	var int_slc []int
	int_slc = s
	sort.Ints(int_slc)
	//fmt.Println("The sorted slice of integers is: ", int_slc)
	c <- int_slc // send sorted array to c
}

func main() {
	c := make(chan []int)
	s := takeInput()
	//sortArray(s, c)
	//z := <-c
	//fmt.Println(z)
	go sortArray(s[:len(s)/4], c)
	go sortArray(s[len(s)/4:len(s)/2], c)
	go sortArray(s[len(s)/2:(3*len(s)/4)], c)
	go sortArray(s[(3*len(s)/4):], c)
	w, x, y, z := <-c, <-c, <-c, <-c // receive from c
	fmt.Println(w, x, y, z)
	//fArray := append(w,x...,y,z...)
	//fmt.Println(fArray)
}
