package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	a := make(map[string]Animal)
	a["cow"] = Animal{"grass", "walk", "moo"}
	a["bird"] = Animal{"worms", "fly", "peep"}
	a["snake"] = Animal{"mice", "slither", "hss"}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		input := strings.Trim(text, " \n")

		s := strings.Split(input, " ")

		if len(s) < 2 {
			fmt.Println("Invalid input. Please enter the animal name and action [Space separated]. Exiting")
			return
		}

		animalName, action := s[0], s[1]

		animal, ok := a[animalName]

		if ok == false {
			fmt.Println("Please give valid animal name. Exiting")
			return
		}

		switch action {
		case "eat":
			animal.Eat()

		case "move":
			animal.Move()

		case "speak":
			animal.Speak()

		default:
			fmt.Println("Enter valid action. Exiting")
		}

	}
}
