package main

var f func(string) int

func test(x string) int {

	return len(x)

}

func main() {

	f = test

}
