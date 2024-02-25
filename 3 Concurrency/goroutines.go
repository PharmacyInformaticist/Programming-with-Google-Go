package main

import (
	"fmt"
	"time"
)

func add_a(val *int) int {
	*val += 1
	fmt.Println(*val)
	return *val
}
func add_b(val *int) int {
	*val += 1
	fmt.Println(*val)
	return *val
}

// Both func will try to read the var num, add 1 to it, and write new value to num
// funcs will read a different value depending on which routine reads num origianl val first
func main() {
	var num int
	num = 1
	go add_a(&num)
	go add_b(&num)
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("Done")
}
