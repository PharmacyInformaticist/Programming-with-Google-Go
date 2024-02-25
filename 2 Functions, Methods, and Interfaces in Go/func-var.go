package main

import "fmt"

var FuncVar func(int) int

func incFn(x int) int {
	return x + 1
}

func main() {
	FuncVar = incFn
	fmt.Println(FuncVar(1))
}
