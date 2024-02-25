package main

import "fmt"

// we can pass a pointer to an array as args for the param
func foo(x *[3]int) int {
	return (*x)[0] = (*x)[0] + 1

}

// array args are copied, this can be a problem for big arrays
func main() {
	a := [3]int{1, 2, 3}
	fmt.Println(foo(a))
}
