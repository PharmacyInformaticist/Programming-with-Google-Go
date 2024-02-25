package main

// https://pkg.go.dev/encoding/json
import (
	"encoding/json"
	"fmt"
	"log"
)

// To encode JSON data we use the Marshal function.
type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000} //and an instance of Message
	b, err := json.Marshal(m)                           // we can marshal a JSON-encoded version of m using json.Marshal:
	if err != nil {
		log.Panic("Panic due to error") // If all is well, err will be nil and b will be a []byte containing this JSON data:
	}
	fmt.Printf("b is of type %T\n", b)
	// b == []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
}
