package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Compares an element of index slice[i] with the at its right
func Swap(slc []int, index int) {
	if slc[index] > slc[index+1] {
		left := slc[index]
		slc[index] = slc[index+1]
		slc[index+1] = left
	}
}

// Sorts the elements modifying the original slice
func BubbleSort(slc *[]int) {
	slc_len := len(*slc)
	for i := 0; i < slc_len; i++ {
		for j := 0; j < slc_len-i-1; j++ {
			Swap(*slc, j)
		}
	}
}

// creates an int slice from user input
func CreateSlice() []int {
	var slc []int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a sequence of up to 10 integers: ")
	scanner.Scan()
	num_input := scanner.Text()
	s := strings.Split(num_input, " ")
	for i := 0; i < len(s); i++ {
		str_num := s[i]
		num, err := strconv.Atoi(str_num)
		if err != nil {
			log.Fatal(err)
		}
		slc = append(slc, num)
	}
	return slc
}

func main() {
	usr_slc := CreateSlice()
	fmt.Println("You entered:\n", usr_slc)
	BubbleSort(&usr_slc)
	fmt.Println("The sorted slice is:\n", usr_slc)
}
