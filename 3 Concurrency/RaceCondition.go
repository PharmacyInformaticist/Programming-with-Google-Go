package main

import (
    "fmt"
    "time"
	"strconv"
)

func pow(array *[5]int){
	for i := 0; i < len(array); i++ {
		if(array[i] % 2 == 0){
			fmt.Println("even")
		}
		array[i] = array[i] * array[i]
	}
}

func printArray(array *[5]int){
	for i := 0; i < len(array); i++ {
		fmt.Printf("%s  ", strconv.Itoa(array[i]))
	}
	println()
}

func main() {
    fmt.Println("hello world")
	var a = [5]int{1, 2, 3, 4, 5}
	go pow(&a)
	go printArray(&a)

	time.Sleep(time.Duration(1)*time.Second)
    fmt.Println("done")
}