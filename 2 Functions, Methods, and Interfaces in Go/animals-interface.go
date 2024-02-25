package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Animal Struc
type myAnimal struct {
	food, locomotion, noise string
}

// Create new animal functions
func newCow() myAnimal {
	var aname = myAnimal{"grass", "walk", "moo"}
	return aname
}

func newBird() myAnimal {
	var aname = myAnimal{"worms", "fly", "peep"}
	return aname
}

func newSnake() myAnimal {
	var aname = myAnimal{"mice", "slither", "hsss"}
	return aname
}

// Interface Methods
func (an myAnimal) Eat() {
	fmt.Println(an.food)
}

func (an myAnimal) Move() {
	fmt.Println(an.locomotion)
}

func (an myAnimal) Speak() {
	fmt.Println(an.noise)
}

// Take user request
func takeRequest() []string {
take_input:
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter request > ")
	scanner.Scan()
	// can take to lower off and add || condition on strings evaluation steps
	data_request := strings.ToLower(scanner.Text())
	req := strings.Split(data_request, " ")
	if len(req) != 3 {
		log.Println("Invalid Input")
		goto take_input
	}
	return req
}

// Map containing existing animals
var animals_map = map[string]myAnimal{}

// Create animal request
func createRequest(atype, aname string, a_map map[string]myAnimal) {
	//call newAnimal() of atype, save aname to Map
	switch atype {
	case "bird":
		bird_name := aname
		aname := newBird()
		a_map[bird_name] = aname
		fmt.Println("created it")
	case "cow":
		cow_name := aname
		aname := newCow()
		a_map[cow_name] = aname
		fmt.Println("created it")
	case "snake":
		snake_name := aname
		aname := newSnake()
		a_map[snake_name] = aname
		fmt.Println("created it")
	default:
		fmt.Println("Cannot create an animal of this type")
	}
}

// Get Interface value
func animalInfo(a Animal, info string) {
	switch info {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Println("wrong action, chose from \"eat\", \"move\", \"speak\"")
	}
}

// Query on animal info request
func queryRequest(aname, anaction string) {
	animal, found := animals_map[aname]
	// retrieve data from interface
	if found {
		animalInfo(animal, anaction)
	} else {
		fmt.Println("wrong animal, chose from an existing from: ", animals_map)
	}
}

// Process user request
func processRequest() {
	usr_req := takeRequest()
	var req, name, type_or_action = usr_req[0], usr_req[1], usr_req[2]
	// check request type: create req or info query
	switch req {
	case "newanimal":
		createRequest(type_or_action, name, animals_map)
	case "query":
		queryRequest(name, type_or_action)
	default:
		fmt.Println("Invalid Request")
	}
}

// main func loop, ask for input until user finishes execution
func main() {
run_main:
	fmt.Println("Initial animals: ", animals_map)
	var ans string = "c"
	for {
		fmt.Println("New animals: ", animals_map)
		processRequest()
		fmt.Println("Press c to continue or press any other key to finish")
		_, err := fmt.Scan(&ans)
		if err != nil {
			log.Println("Invalid Input")
			goto run_main
		}
		if ans != "c" {
			fmt.Println("Program terminated")
			break
		}
	}
}
