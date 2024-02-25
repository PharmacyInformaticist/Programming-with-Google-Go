package main

import "fmt"

// ellipsis ... implies any number of int inputs
func GetMax(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func main() {
	fmt.Println(GetMax(1, 3, 6, 4))
	// can pass an slice to a variadic func
	// ints prepackaged into a slice passed as args
	vslice := []int{1, 3, 6, 4}
	fmt.Println(GetMax(vslice...))
}
