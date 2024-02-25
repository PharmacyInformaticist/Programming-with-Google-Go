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

type myAnimal struct {
	food, locomotion, noise string
}

// create animal
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

// Methods
// Food eaten
func (an myAnimal) Eat() {
	fmt.Println(an.food)
}

// Locomotion method
func (an myAnimal) Move() {
	fmt.Println(an.locomotion)
}

// Spoken sound
func (an myAnimal) Speak() {
	fmt.Println(an.noise)
}

// Take user request
func takeRequest() []string {
take_input:
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter request > ")
	scanner.Scan()
	data_request := strings.ToLower(scanner.Text())
	req := strings.Split(data_request, " ")
	if len(req) != 3 {
		log.Println("Invalid Input")
		goto take_input
	}
	return req
}

// existing animals
var animals_map = map[string]myAnimal{}

//animals_map = make(map[string]myAnimal)

// Process user request
func createRequest(atype, aname string, a_map map[string]myAnimal) {

	//animals_map = make(map[string]myAnimal)
	//call newAnimal(), save var to name
	// create map of newanimals
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
	}
}
func queryRequest(aname, anaction string) {
	animal, found := animals_map[aname]
	if found {
		if anaction == "eat" {
			animal.Eat()
		} else if anaction == "move" {
			animal.Move()
		} else if anaction == "speak" {
			animal.Speak()
		} else {
			fmt.Println("wrong action, chose from \"eat\", \"move\", \"speak\"")
		}
	} else {
		fmt.Println("wrong animal, chose from an existing \"cow\", \"bird\", \"snake\"")
	}
}

func processRequest() { //amap *map[string]myAnimal
	usr_req := takeRequest()
	var req, name, type_or_action = usr_req[0], usr_req[1], usr_req[2]
	// animals_map = make(map[string]myAnimal)
	// check if create req or info query
	switch req {
	case "newanimal":
		//call newAnimal(), save var to name
		// create map of newanimals
		createRequest(type_or_action, name, animals_map)
	case "query":
		// search for existing animals and call methods
		queryRequest(name, type_or_action)
	}
}

// main func loops ask for input until user finishes execution
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
