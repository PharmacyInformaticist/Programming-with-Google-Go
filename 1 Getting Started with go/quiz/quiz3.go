package main

import "fmt"

func main() {
	x := [...]int{1, 2, 3, 4, 5}
	y := x[0:2] // [1,2]
	z := x[1:4] // [2,3,4] remenber cap diminishes
	fmt.Print(len(y), cap(y), len(z), cap(z))
}
