package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test(value *int, funDelay func()) {
	funDelay()
	if *value == 0 {
		*value = 1
		funDelay()
	} else {
		*value = 2
	}
	fmt.Printf("%d\n", *value)
}

func main() {
	var fun = func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Microsecond)
	}
	var fun2 = func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Microsecond)
	}
	var a = 0
	go test(&a, fun)
	go test(&a, fun2)

	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("done")
}
