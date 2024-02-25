package main

// passing a slice passes a pointer to the underlying array
import "fmt"

func foo(sli []int) []int {
	return (sli[0] = sli[0] + 1)
}

func main() {
	a := []int{1, 2, 3}
	foo(a)
	fmt.Print(a)
}
