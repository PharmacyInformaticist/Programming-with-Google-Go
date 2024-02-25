package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	m := make(map[string]string)
	var (
		n, a string
	)
	fmt.Println("Please enter a name: ")
	fmt.Scan(&n)
	m["name"] = n
	fmt.Println("Please enter an address: ")
	fmt.Scan(&a)
	m["address"] = a
	barr, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The JSON object is: ", barr)
	fmt.Println("The JSON object content is: ", string(barr))
}
