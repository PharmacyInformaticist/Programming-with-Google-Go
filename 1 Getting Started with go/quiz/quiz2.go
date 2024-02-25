package main

import "fmt"

func main() {
	x := [...]int{4, 8, 5} //[1,8,3]
	y := x[0:2]            // [0,8][1,8]
	z := x[1:3]            // [8,5][8,3]
	y[0] = 1
	z[1] = 3
	fmt.Print(x)
}
