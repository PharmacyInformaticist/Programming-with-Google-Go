package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.WriteFile("outfile.txt", []byte("This is my first file written in Golang :) "), 0777)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("done")
	// an alternative way
	// file, err := os.Create("outfile.txt")
	// nfil, err2 := file.WriteString("Hello World")
}
