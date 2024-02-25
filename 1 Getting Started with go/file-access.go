package main

// https://pkg.go.dev/os
// https://pkg.go.dev/io/ioutil
import (
	"fmt" // io/ioutil deprecated use os.readfile
	"log"
	"os"
)

func main() {
	dat, e := os.ReadFile("words.txt")
	// dat is []bye filled with the contents of the entire file
	if e != nil {
		log.Fatal(e)
	}
	fmt.Printf("dat is of the type %T\n", dat)
	fmt.Println(string(dat))
}
