package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Main Type
type Animal struct {
	food, locomotion, noise string
}

// animals
var (
	cow   = Animal{"grass", "walk", "moo"}
	bird  = Animal{"worms", "fly", "peep"}
	snake = Animal{"mice", "slither", "hsss"}
)

// Methods
// Food eaten
func (an Animal) Eat() {
	fmt.Println(an.food)
}

// Locomotion method
func (an Animal) Move() {
	fmt.Println(an.locomotion)
}

// Spoken sound
func (an Animal) Speak() {
	fmt.Println(an.noise)
}

// Take user request
func takeRequest() []string {
take_input:
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter > animal information-requested")
	scanner.Scan()
	data_request := strings.ToLower(scanner.Text())
	req := strings.Split(data_request, " ")
	if len(req) != 2 {
		log.Println("Invalid Input")
		goto take_input
	}
	return req
}

// Process user request
func processRequest() {
user_input:
	req := takeRequest()
	switch an := req[0]; an {
	case "bird":
		switch info := req[1]; info {
		case "eat":
			fmt.Println(bird.food)
		case "move":
			fmt.Println(bird.locomotion)
		case "speak":
			fmt.Println(bird.noise)
		default:
			log.Printf("%s about a %s is not available, please try again", info, an)
			goto user_input
		}
	case "cow":
		switch info := req[1]; info {
		case "eat":
			fmt.Println(cow.food)
		case "move":
			fmt.Println(cow.locomotion)
		case "speak":
			fmt.Println(cow.noise)
		default:
			log.Printf("%s about a %s is not available, please try again", info, an)
			goto user_input
		}
	case "snake":
		switch info := req[1]; info {
		case "eat":
			fmt.Println(snake.food)
		case "move":
			fmt.Println(snake.locomotion)
		case "speak":
			fmt.Println(snake.noise)
		default:
			log.Printf("%s about a %s is not available, please try again", info, an)
			goto user_input
		}
	default:
		log.Printf("There is no information about %s, please try again", an)
		goto user_input
	}
}

// main func loops ask for input until user finishes execution
func main() {
run_main:
	var ans string = "c"
	for {
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
