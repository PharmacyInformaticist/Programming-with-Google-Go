// python lambda
package main

import "fmt"

func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

func main() {
	v := applyIt(func(x int) int { return x + 1 }, 2)
	fmt.Println(v)
}
