package main

// https://go.dev/blog/json
import (
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m1 := Message{"Alice", "Hello", 1294706395881547000}
	var m2 Message
	b, err := json.Marshal(m1) // we can marshal a JSON-encoded version of m using json.Marshal:
	if err != nil {
		log.Panic("Panic due to error") // If all is well, err will be nil and b will be a []byte containing this JSON data:
	}
	fmt.Printf("b is of type %T\n", b)
	err2 := json.Unmarshal(b, &m2) // Unmarshal will decode only the fields that it can find in the destination type
	if err2 != nil {
		log.Panic("Panic due to error") // If all is well, err will be nil and b will be a []byte containing this JSON data:
	}
	fmt.Printf("b is of type %T\n", m2)
	fmt.Println(m2)
}
