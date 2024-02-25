package main

import (
	"fmt"
	"time"
)

func test(value *int) {
	if *value == 0 {
		*value = 1
	} else {
		*value = 2
	}
	fmt.Printf("%d\n", *value)
}

func main() {
	var a = 0
	go test(&a)
	go test(&a)

	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("done")
}
